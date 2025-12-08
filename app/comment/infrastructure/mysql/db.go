package mysql

import (
	"context"
	"errors"
	"myreel/app/comment/domain/model"
	"myreel/app/comment/domain/repository"
	"myreel/pkg/errno"
	"time"

	"gorm.io/gorm"
)

type commentDB struct {
	client *gorm.DB
}

func NewCommentDB(client *gorm.DB) repository.CommentDB {
	return &commentDB{client: client}
}

func (db *commentDB) Magrate() error {
	if err := db.client.AutoMigrate(&Comment{}); err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "mysql: failed to auto magrate comment model")
	}
	return nil
}

func (db *commentDB) CreateComment(ctx context.Context, comment *model.Comment) error {
	c := &Comment{
		Id:       comment.Id,
		Uid:      comment.Uid,
		VideoId:  comment.VideoId,
		ParentId: comment.ParentId,
		Content:  comment.Content,
	}
	if err := db.client.WithContext(ctx).Create(c).Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "failed to create comment")
	}
	return nil
}

func (db *commentDB) AddChildCount(ctx context.Context, commentId int64) error {
	if err := db.client.WithContext(ctx).
		Model(&Comment{}).
		Where("id = ?", commentId).
		UpdateColumn("child_count", gorm.Expr("child_count + 1")).
		Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "failed to add comment child count: %v", err)
	}
	return nil
}

func (db *commentDB) SubtractChildCount(ctx context.Context, commentId int64) error {
	if err := db.client.WithContext(ctx).
		Model(&Comment{}).
		Where("id = ?", commentId).
		UpdateColumn("child_count", gorm.Expr("GREATEST(child_count - 1, 0)")).
		Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "failed to add comment child count: %v", err)
	}
	return nil
}

func (db *commentDB) GetCommentListByVideoId(ctx context.Context, videoId, cursor, limit int64) ([]*model.Comment, int64, error) {
	var comments []*Comment
	var total int64
	tx := db.client.WithContext(ctx).
		Model(&Comment{}).
		Where("video_id = ?", videoId).
		Where("parent_id = 0").
		Limit(int(limit)).
		Order("created_at DESC")

	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get comments count by video id: %v", err)
	}

	if cursor != 0 {
		tx = tx.Where("created_at < ?", time.Unix(cursor, 0))
	}

	if err := tx.Find(&comments).
		Error; err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get comments by video id: %v", err)
	}

	l := len(comments)
	result := make([]*model.Comment, l)
	if l > 0 {
		for i, comment := range comments {
			result[i] = &model.Comment{
				Id:         comment.Id,
				VideoId:    comment.VideoId,
				Uid:        comment.Uid,
				ParentId:   comment.ParentId,
				LikeCount:  comment.LikeCount,
				ChildCount: comment.ChildCount,
				Content:    comment.Content,
				CreatedAt:  comment.CreatedAt.Unix(),
				UpdatedAt:  comment.UpdatedAt.Unix(),
			}
		}
	}
	return result, total, nil
}

func (db *commentDB) GetCommentListByCommentId(ctx context.Context, commentId, cursor, limit int64) ([]*model.Comment, int64, error) {
	var comments []*Comment
	var total int64
	tx := db.client.WithContext(ctx).
		Model(&Comment{}).
		Where("parent_id = ?", commentId).
		Limit(int(limit)).
		Order("created_at DESC")

	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get comments count by comment id: %v", err)
	}

	if cursor != 0 {
		tx = tx.Where("created_at < ?", time.Unix(cursor, 0))
	}

	if err := tx.Find(&comments).
		Error; err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get comments by comment id: %v", err)
	}

	l := len(comments)
	result := make([]*model.Comment, l)
	if l > 0 {
		for i, comment := range comments {
			result[i] = &model.Comment{
				Id:         comment.Id,
				VideoId:    comment.VideoId,
				Uid:        comment.Uid,
				ParentId:   comment.ParentId,
				LikeCount:  comment.LikeCount,
				ChildCount: comment.ChildCount,
				Content:    comment.Content,
				CreatedAt:  comment.CreatedAt.Unix(),
				UpdatedAt:  comment.UpdatedAt.Unix(),
			}
		}
	}
	return result, total, nil
}

func (db *commentDB) DeleteCommentById(ctx context.Context, id int64) error {
	if err := db.client.WithContext(ctx).
		Delete(&Comment{}, id).
		Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to delete comment by id: %v", err)
	}
	return nil
}

func (db *commentDB) DeleteCommentsByVideoId(ctx context.Context, id int64) error {
	if err := db.client.WithContext(ctx).
		Where("video_id = ?", id).
		Delete(&Comment{}).
		Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to delete comments by video id: %v", err)
	}
	return nil
}

func (db *commentDB) GetCommentById(ctx context.Context, id int64) (*model.Comment, error) {
	var c Comment
	if err := db.client.WithContext(ctx).
		First(&c, id).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.CommentNotFound
		}
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "failed to get comment by id: %v", err)
	}
	return &model.Comment{
		Id:         c.Id,
		VideoId:    c.VideoId,
		Uid:        c.Uid,
		ParentId:   c.ParentId,
		LikeCount:  c.LikeCount,
		ChildCount: c.ChildCount,
		Content:    c.Content,
		CreatedAt:  c.CreatedAt.Unix(),
		UpdatedAt:  c.UpdatedAt.Unix(),
	}, nil
}

func (db *commentDB) AddLikeCount(ctx context.Context, id int64) error {
	if err := db.client.WithContext(ctx).
		Model(&Comment{}).
		Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count + ?", 1)).Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to add comment's like count: %v", err)
	}
	return nil
}

func (db *commentDB) SubtractLikeCount(ctx context.Context, id int64) error {
	if err := db.client.WithContext(ctx).
		Model(&Comment{}).
		Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count - ?", 1)).Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to subtract comment's like count: %v", err)
	}
	return nil
}

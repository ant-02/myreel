package mysql

import (
	"context"
	"myreel/app/comment/domain/model"
	"myreel/app/comment/domain/repository"
	"myreel/pkg/errno"

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

// func (cr *commentRepository) GetCommentListByVideoId(videoId string, pageNum, pageSize int64) ([]*model.Comment, error) {
// 	var comments []*model.Comment
// 	err := cr.db.Where("video_id = ?", videoId).
// 		Where("parent_id IS NULL").
// 		Offset((int(pageNum) - 1) * int(pageSize)).
// 		Limit(int(pageSize)).
// 		Find(&comments).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return comments, nil
// }

// func (cr *commentRepository) GetCommentListByCommentId(commentId string, pageNum, pageSize int64) ([]*model.Comment, error) {
// 	var comments []*model.Comment
// 	err := cr.db.Where("parent_id = ?", commentId).
// 		Offset((int(pageNum) - 1) * int(pageSize)).
// 		Limit(int(pageSize)).
// 		Find(&comments).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return comments, nil
// }

// func (cr *commentRepository) DeleteCommentsByVideoId(videoId string) error {
// 	return cr.db.Where("video_id = ?", videoId).
// 		Delete(&model.Comment{}).Error
// }

// func (cr *commentRepository) DeleteCommentById(id string) error {
// 	return cr.db.Delete(&model.Comment{}, id).Error
// }

package mysql

import (
	"context"
	"errors"
	"myreel/app/like/domain/model"
	"myreel/app/like/domain/repository"
	"myreel/pkg/errno"

	"gorm.io/gorm"
)

type likeDB struct {
	client *gorm.DB
}

func NewLikeDB(client *gorm.DB) repository.LikeDB {
	return &likeDB{client: client}
}

func (db *likeDB) Magrate() error {
	if err := db.client.AutoMigrate(&Like{}); err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "mysql: failed to auto magrate like model")
	}
	return nil
}

func (db *likeDB) GetVideoLike(ctx context.Context, videoId, uid int64) (*model.Like, error) {
	var like model.Like
	err := db.client.WithContext(ctx).
		Where("uid = ? and video_id = ?", uid, videoId).
		Where("deleted_at IS NULL").
		First(&like).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.LikeNotFound
		}
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get video like: %v", err)
	}
	return &like, nil
}

func (db *likeDB) GetCommentLike(ctx context.Context, commentId, uid int64) (*model.Like, error) {
	var like model.Like
	err := db.client.WithContext(ctx).
		Where("uid = ? and comment_id = ?", uid, commentId).
		Where("deleted_at IS NULL").
		First(&like).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.LikeNotFound
		}
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get video like: %v", err)
	}
	return &like, nil
}

func (db *likeDB) SetLikeStatus(ctx context.Context, id int64, status int64) error {
	err := db.client.WithContext(ctx).
		Model(&Like{}).
		Where("id = ?", id).
		Where("deleted_at IS NULL").
		Update("status", status).Error
	if err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to set like status: %v", err)
	}
	return nil
}

func (db *likeDB) CreateLike(l *model.Like) error {
	dl := &Like{
		Id:        l.Id,
		Uid:       l.Uid,
		CommentId: l.CommentId,
		VideoId:   l.VideoId,
	}
	err := db.client.Create(&dl).Error
	if err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "failed to create like: %v", err)
	}
	return nil
}

// func (db *likeDB) GetVideoLikeList(uid string, pageNum, pageSize int64) ([]*string, error) {
// 	var videoIds []*string
// 	err := lr.db.Model(&model.Like{}).
// 		Select("video_id").
// 		Where("uid = ?", uid).
// 		Where("video_id IS NOT NULL").
// 		Where("status = 1").
// 		Where("deleted_at IS NULL").
// 		Offset((int(pageNum) - 1) * int(pageSize)).
// 		Limit(int(pageSize)).
// 		Find(&videoIds).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return videoIds, nil
// }

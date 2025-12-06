package mysql

import (
	"context"
	"myreel/app/video/domain/model"
	"myreel/app/video/domain/repository"
	"myreel/pkg/errno"
	"time"

	"gorm.io/gorm"
)

type videoDB struct {
	client *gorm.DB
}

func NewVideoDB(client *gorm.DB) repository.VideoDB {
	return &videoDB{client: client}
}

func (db *videoDB) Magrate() error {
	if err := db.client.AutoMigrate(&Video{}); err != nil {
		return errno.NewErrNo(errno.InternalDatabaseErrorCode, "mysql: failed to auto magrate video model")
	}
	return nil
}

func (db *videoDB) GetVideosByLatestTime(ctx context.Context, latestTime time.Time) ([]*model.Video, error) {
	var videos []*Video

	err := db.client.WithContext(ctx).
		Where("created_at > ?").
		Where("deleted_at IS NULL").
		Find(&videos).Error
	if err != nil {
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get videos by latest time: %v", err)
	}

	l := len(videos)
	if l == 0 {
		return nil, nil
	}

	result := make([]*model.Video, l)
	for i, v := range videos {
		result[i] = &model.Video{
			Id:           v.Id,
			Uid:          v.Uid,
			Title:        v.Title,
			Description:  v.Description,
			VideoUrl:     v.VideoUrl,
			CoverUrl:     v.CoverUrl,
			VisitCount:   v.VisitCount,
			LikeCount:    v.LikeCount,
			CommentCount: v.CommentCount,
			CreatedAt:    v.CreatedAt.Unix(),
			UpdatedAt:    v.UpdatedAt.Unix(),
		}
	}

	return result, nil
}

func (db *videoDB) CreateVideo(ctx context.Context, video *model.Video) error {
	v := &Video{
		Id:          video.Id,
		Uid:         video.Uid,
		Title:       video.Title,
		Description: video.Description,
		VideoUrl:    video.VideoUrl,
		CoverUrl:    video.CoverUrl,
	}
	err := db.client.WithContext(ctx).Create(v).Error
	if err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to create video : %v", err)
	}
	return nil
}

func (db *videoDB) GetVideosByUid(ctx context.Context, uid, cursor, limit int64) ([]*model.Video, int64, error) {
	var videos []*Video
	var total int64
	var err error

	tx := db.client.WithContext(ctx).Model(&Video{}).
		Where("uid = ?", uid).
		Where("deleted_at IS NULL")

	err = tx.Count(&total).Error
	if err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get video count by user id: %v", err)
	}

	if cursor != 0 {
		tx.Where("id < ?", cursor)
	}

	err = tx.Limit(int(limit)).
		Order("id DESC").
		Find(&videos).Error
	if err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get videos by user id: %v", err)
	}

	l := len(videos)
	result := make([]*model.Video, l)
	if l > 0 {
		for i, v := range videos {
			result[i] = &model.Video{
				Id:           v.Id,
				Uid:          v.Uid,
				Title:        v.Title,
				Description:  v.Description,
				VideoUrl:     v.VideoUrl,
				CoverUrl:     v.CoverUrl,
				VisitCount:   v.VisitCount,
				LikeCount:    v.VisitCount,
				CommentCount: v.CommentCount,
				CreatedAt:    v.CreatedAt.Unix(),
				UpdatedAt:    v.UpdatedAt.Unix(),
			}
		}
	}

	return result, total, nil
}

func (db *videoDB) GetVideosByKeywords(ctx context.Context, keywords string, fromDate, toDate, cursor time.Time, uid, limit int64) ([]*model.Video, int64, error) {
	var videos []*Video
	var total int64
	var err error

	tx := db.client.WithContext(ctx).
		Model(&Video{}).
		Where("deleted_at IS NULL").
		Where("title LIKE ? or description LIKE ?", "%"+keywords+"%", "%"+keywords+"%").
		Where("created_at between ? and ?", fromDate, toDate)

	if uid != 0 {
		tx = tx.Where("uid = ?", uid)
	}

	err = tx.Count(&total).Error
	if err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get count by search: %v", err)
	}

	err = tx.Where("created_at < ?", cursor).
		Limit(int(limit)).
		Order("created_at DESC").
		Find(&videos).Error

	if err != nil {
		return nil, 0, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get videos by search: %v", err)
	}

	l := len(videos)
	result := make([]*model.Video, l)
	if l > 0 {
		for i, v := range videos {
			result[i] = &model.Video{
				Id:           v.Id,
				Uid:          v.Uid,
				Title:        v.Title,
				Description:  v.Description,
				VideoUrl:     v.VideoUrl,
				CoverUrl:     v.CoverUrl,
				VisitCount:   v.VisitCount,
				LikeCount:    v.VisitCount,
				CommentCount: v.CommentCount,
				CreatedAt:    v.CreatedAt.Unix(),
				UpdatedAt:    v.UpdatedAt.Unix(),
			}
		}
	}

	return result, total, nil
}

func (db *videoDB) GetVideoById(ctx context.Context, id int64) (*model.Video, error) {
	var video model.Video
	if err := db.client.WithContext(ctx).
		Where("deleted_at IS NULL").
		First(&video, id).Error; err != nil {
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get video by id: %v", err)
	}
	return &video, nil
}

func (db *videoDB) AddLikeCount(ctx context.Context, id int64) error {
	if err := db.client.WithContext(ctx).
		Model(&Video{}).
		Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count + ?", 1)).Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to add like_count: %v", err)
	}
	return nil
}

func (db *videoDB) SubtractLikeCount(ctx context.Context, id int64) error {
	if err := db.client.WithContext(ctx).
		Model(&Video{}).
		Where("id = ?", id).
		Update("like_count", gorm.Expr("like_count - ?", 1)).Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to subtract like_count: %v", err)
	}
	return nil
}

func (db *videoDB) AddCommentCount(ctx context.Context, id int64) error {
	if err := db.client.WithContext(ctx).
		Model(&Video{}).
		Where("id = ?", id).
		Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
		return errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to add comment_count: %v", err)
	}
	return nil
}

// func (db *videoDB) GetVideosByIds(ids []*string) ([]*Video, error) {
// 	var videos []*Video
// 	err := db.client.Where("id IN ?", ids).Find(&videos).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return videos, nil
// }

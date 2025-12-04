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

func (db *videoDB) GetVideosByLatestTime(ctx context.Context, latestTime int64) ([]*model.Video, error) {
	var videos []*Video

	err := db.client.WithContext(ctx).
		Where("created_at > ?", time.Unix(latestTime, 0)).
		Where("deleted_at IS NULL").
		Find(&videos).Error
	if err != nil {
		return nil, errno.Errorf(errno.InternalDatabaseErrorCode, "mysql: failed to get videos by latest time: %v", err)
	}

	l := len(videos)
	if l == 0 {
		return nil, nil
	}

	result := make([]*model.Video, 0, l)
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

// func (db *videoDB) GetVideosByUid(uid string, pageNum, pageSize int64) ([]*Video, int64, error) {
// 	var videos []*Video
// 	var total int64
// 	var err error

// 	tx := db.client.Model(&Video{}).
// 		Where("uid = ?", uid).
// 		Where("deleted_at IS NULL")

// 	err = tx.Count(&total).Error
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	err = db.client.Where("uid = ?", uid).
// 		Where("deleted_at IS NULL").
// 		Offset((int(pageNum) - 1) * int(pageSize)).
// 		Limit(int(pageSize)).
// 		Find(&videos).Error
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	return videos, total, nil
// }

// func (db *videoDB) GetVideosGroupByVisitCount(pageNum, pageSize int64) ([]*Video, error) {
// 	var videos []*Video
// 	first, end := (pageNum-1)*pageSize, pageNum*pageSize
// 	instance := database.GetRedisInstance()
// 	ctx := context.Background()
// 	videoStrings, err := instance.LRange(ctx, key, first, end)
// 	if err != nil || videoStrings == nil {
// 		return nil, err
// 	}
// 	if len(videoStrings) > 0 {
// 		for _, s := range videoStrings {
// 			var v Video
// 			if err := json.Unmarshal([]byte(s), &v); err != nil {
// 				return nil, err
// 			}
// 			videos = append(videos, &v)
// 		}
// 		return videos, nil
// 	}

// 	err = db.client.Where("deleted_at IS NULL").
// 		Order("visit_count desc").
// 		Offset(int(first)).
// 		Limit(int(pageSize)).
// 		Find(&videos).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	for _, v := range videos {
// 		j, err := json.Marshal(v)
// 		if err != nil {
// 			return nil, err
// 		}
// 		instance := database.GetRedisInstance()
// 		ctx := context.Background()
// 		err = instance.RPush(ctx, key, j)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return videos, nil
// }

// func (db *videoDB) GetVideosByKeywords(keywords, fromDate, toDate, uid string, pageNum, pageSize int64) ([]*Video, int64, error) {
// 	var videos []*Video
// 	var total int64
// 	var err error
// 	tx := db.client.Model(&Video{}).
// 		Where("title LIKE ? or description LIKE ?", "%"+keywords+"%", "%"+keywords+"%")

// 	if fromDate != "" {
// 		tx = tx.Where("from_date > ?", fromDate)
// 	}
// 	if toDate != "" {
// 		tx = tx.Where("to_date < ?", toDate)
// 	}
// 	if uid != "" {
// 		tx = tx.Where("uid = ?", uid)
// 	}

// 	err = tx.Count(&total).Error
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	err = tx.Offset((int(pageNum) - 1) * int(pageSize)).
// 		Limit(int(pageSize)).
// 		Find(&videos).Error

// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	return videos, total, nil
// }

// func (db *videoDB) AddLikeCount(id string) error {
// 	return db.client.Model(&Video{}).Where("id = ?", id).Update("like_count", gorm.Expr("like_count + ?", 1)).Error
// }

// func (db *videoDB) SubtractLikeCount(id string) error {
// 	return db.client.Model(&Video{}).Where("id = ?", id).Update("like_count", gorm.Expr("like_count - ?", 1)).Error
// }

// func (db *videoDB) GetVideosByIds(ids []*string) ([]*Video, error) {
// 	var videos []*Video
// 	err := db.client.Where("id IN ?", ids).Find(&videos).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return videos, nil
// }

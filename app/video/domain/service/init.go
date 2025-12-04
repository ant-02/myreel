package service

import (
	"context"
	"myreel/app/video/domain/model"
	"myreel/app/video/domain/repository"
	"myreel/pkg/upyun"
	"myreel/pkg/util"
)

type videoService struct {
	db    repository.VideoDB
	cache repository.VideoCache
	sf    *util.Snowflake
}

type VideoService interface {
	GetVideosByLatestTime(ctx context.Context, latestTime string) ([]*model.Video, error)
	GetVideoUploadToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error)
	GetVideoCoverUploadToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error)
	GenerateVideoId() (int64, error)
	SaveVideo(ctx context.Context, video *model.Video) error
	GetVideosByUserId(ctx context.Context, uid, cursor, limit int64) ([]*model.Video, *model.Pagination, error)
}

func NewVideoService(db repository.VideoDB, sf *util.Snowflake, cache repository.VideoCache) VideoService {
	if db == nil {
		panic("videoService`s db should not be nil")
	}

	svc := &videoService{
		db:    db,
		sf:    sf,
		cache: cache,
	}
	return svc
}

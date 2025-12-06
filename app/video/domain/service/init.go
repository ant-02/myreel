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
	vRpc  repository.RpcPort
}

type VideoService interface {
	GetVideosByLatestTime(ctx context.Context, latestTime string) ([]*model.Video, error)
	GetVideoUploadToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error)
	GetVideoCoverUploadToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error)
	GenerateVideoId() (int64, error)
	SaveVideo(ctx context.Context, video *model.Video) error
	GetVideosByUserId(ctx context.Context, uid, cursor, limit int64) ([]*model.Video, *model.Pagination, error)
	GetVideosByKeywords(ctx context.Context, keywords string, fromDate, toDate, cursor, uid, limit int64) ([]*model.Video, *model.Pagination, error)
	GetVideosByIds(ctx context.Context, ids []int64) ([]*model.Video, error)
	CalculateHotScore(video *model.Video) int64
	DecrVideoLike(ctx context.Context, videoId int64) error
	IncrVideoLike(ctx context.Context, videoId int64) error
	AddCommentCount(ctx context.Context, id int64) error
}

func NewVideoService(db repository.VideoDB, sf *util.Snowflake, cache repository.VideoCache, vRpc repository.RpcPort) VideoService {
	if db == nil {
		panic("videoService`s db should not be nil")
	}

	svc := &videoService{
		db:    db,
		sf:    sf,
		cache: cache,
		vRpc:  vRpc,
	}
	return svc
}

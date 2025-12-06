package usecase

import (
	"context"
	"myreel/app/video/domain/model"
	"myreel/app/video/domain/repository"
	"myreel/app/video/domain/service"
	"myreel/pkg/upyun"
)

type useCase struct {
	db    repository.VideoDB
	svc   service.VideoService
	cache repository.VideoCache
	vRpc  repository.RpcPort
}

type VideoUseCase interface {
	GetVideosByLatestTime(ctx context.Context, latestTime string) ([]*model.Video, error)
	GetVideoUplaodToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error)
	GetVideoCoverUplaodToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error)
	SaveVideo(ctx context.Context, video *model.Video) error
	GetVideosByUserId(ctx context.Context, uid, cursor, limit int64) ([]*model.Video, *model.Pagination, error)
	GetVideosGroupByVisitCount(ctx context.Context, cursor, limit int64) ([]*model.Video, *model.Pagination, error)
	GetVideosByKeywords(ctx context.Context, keywords, userrname string, fromDate, toDate, cursor, limit int64) ([]*model.Video, *model.Pagination, error)
	VideoLikeAction(ctx context.Context, videoId, actionType int64) error
	GetVideosByIds(ctx context.Context, ids []int64) ([]*model.Video, error)
	AddCommentCount(ctx context.Context, id int64) error
}

func NewVideoUseCase(db repository.VideoDB, svc service.VideoService, cache repository.VideoCache, vRpc repository.RpcPort) VideoUseCase {
	return &useCase{
		db:    db,
		svc:   svc,
		cache: cache,
		vRpc:  vRpc,
	}
}

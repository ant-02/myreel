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
}

type VideoUseCase interface {
	GetVideosByLatestTime(ctx context.Context, latestTime string) ([]*model.Video, error)
	GetVideoUplaodToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error)
}

func NewVideoUseCase(db repository.VideoDB, svc service.VideoService, cache repository.VideoCache) *useCase {
	return &useCase{
		db:    db,
		svc:   svc,
		cache: cache,
	}
}

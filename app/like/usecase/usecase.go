package usecase

import (
	"context"
	"myreel/app/like/domain/repository"
	"myreel/app/like/domain/service"
)

type useCase struct {
	db    repository.LikeDB
	svc   service.LikeService
	cache repository.LikeCache
	lRpc  repository.RpcPort
}

type LikeUseCase interface {
	LikeAction(ctx context.Context, videoId, commentId, uid, actionType int64) error
}

func NewLikeUseCase(db repository.LikeDB, svc service.LikeService, cache repository.LikeCache, lRpc repository.RpcPort) *useCase {
	return &useCase{
		db:    db,
		svc:   svc,
		cache: cache,
		lRpc:  lRpc,
	}
}

package service

import (
	"context"
	"myreel/app/like/domain/model"
	"myreel/app/like/domain/repository"
	"myreel/pkg/util"
)

type likeService struct {
	db   repository.LikeDB
	sf   *util.Snowflake
	lRpc repository.RpcPort
}

type LikeService interface {
	GenerateLikeId() (int64, error)
	CreateLike(ctx context.Context, l *model.Like) error
}

func NewLikeService(db repository.LikeDB, sf *util.Snowflake, lRpc repository.RpcPort) LikeService {
	if db == nil {
		panic("LikeService`s db should not be nil")
	}

	svc := &likeService{
		db:   db,
		sf:   sf,
		lRpc: lRpc,
	}
	return svc
}

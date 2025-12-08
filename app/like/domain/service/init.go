package service

import (
	"context"
	"myreel/app/like/domain/model"
	"myreel/app/like/domain/repository"
	"myreel/pkg/util"
)

type likeService struct {
	db    repository.LikeDB
	sf    *util.Snowflake
	cache repository.LikeCache
	lRpc  repository.RpcPort
}

type LikeService interface {
	GenerateLikeId() (int64, error)
	CreateLike(ctx context.Context, l *model.Like) error
	VideoUserLikeAction(ctx context.Context, videoId, uid, actionType int64) error
	CommentUserLikeAction(ctx context.Context, commentId, uid, actionType int64) error
}

func NewLikeService(db repository.LikeDB, sf *util.Snowflake, cache repository.LikeCache, lRpc repository.RpcPort) LikeService {
	if db == nil {
		panic("LikeService`s db should not be nil")
	}

	svc := &likeService{
		db:    db,
		sf:    sf,
		cache: cache,
		lRpc:  lRpc,
	}
	return svc
}

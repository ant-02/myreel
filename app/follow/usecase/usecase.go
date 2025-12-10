package usecase

import (
	"context"
	"myreel/app/follow/domain/model"
	"myreel/app/follow/domain/repository"
	"myreel/app/follow/domain/service"
)

type useCase struct {
	db  repository.FollowDB
	svc service.FollowService
	rpc repository.RpcPort
}

type FollowUseCase interface {
	FollowAction(ctx context.Context, userId, toUserId, actionType int64) error
	GetUsersByFolloweredId(ctx context.Context, userId, cursor, limit int64) ([]*model.UserProfile, *model.Pagination, error)
	GetUsersByFolloweringId(ctx context.Context, userId, cursor, limit int64) ([]*model.UserProfile, *model.Pagination, error)
}

func NewFollowUseCase(db repository.FollowDB, svc service.FollowService, rpc repository.RpcPort) FollowUseCase {
	return &useCase{
		db:  db,
		svc: svc,
		rpc: rpc,
	}
}

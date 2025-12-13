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
	GetFriendsById(ctx context.Context, id, cursor, limit int64) ([]*model.UserProfile, *model.Pagination, error)
	CreateGroup(ctx context.Context, creatorId int64, name string, userIds ...int64) error
	GetGroupByJoined(ctx context.Context, userId int64) ([]*model.Group, error)
	GetGroupByCreator(ctx context.Context, creatorId int64) ([]*model.Group, error)
}

func NewFollowUseCase(db repository.FollowDB, svc service.FollowService, rpc repository.RpcPort) FollowUseCase {
	return &useCase{
		db:  db,
		svc: svc,
		rpc: rpc,
	}
}

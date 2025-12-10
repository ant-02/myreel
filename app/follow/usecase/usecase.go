package usecase

import (
	"context"
	"myreel/app/follow/domain/repository"
	"myreel/app/follow/domain/service"
)

type useCase struct {
	db  repository.FollowDB
	svc service.FollowService
}

type FollowUseCase interface {
	FollowAction(ctx context.Context, userId, toUserId, actionType int64) error
}

func NewFollowUseCase(db repository.FollowDB, svc service.FollowService) FollowUseCase {
	return &useCase{
		db:  db,
		svc: svc,
	}
}

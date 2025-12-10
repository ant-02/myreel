package service

import (
	"context"
	"myreel/app/follow/domain/model"
	"myreel/app/follow/domain/repository"
	"myreel/pkg/util"
)

type FollowService interface {
	FollowAction(ctx context.Context, userId, toUserId, actionType int64) error
	CreateFollow(ctx context.Context, f *model.Follow) error
	GenerateFollowId() (int64, error)
}

type followService struct {
	db repository.FollowDB
	sf *util.Snowflake
}

func NewFollowService(db repository.FollowDB, sf *util.Snowflake) FollowService {
	if db == nil {
		panic("followService`s db should not be nil")
	}

	svc := &followService{
		db: db,
		sf: sf,
	}
	return svc
}

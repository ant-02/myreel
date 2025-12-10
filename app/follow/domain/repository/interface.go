package repository

import (
	"context"
	"myreel/app/follow/domain/model"
	"time"
)

type FollowDB interface {
	Magrate() error
	GetFollowByUserIdAndToUserId(ctx context.Context, userId, toUserId int64) (*model.Follow, error)
	SetFollowStatus(ctx context.Context, id, status int64) error
	CreateFollow(ctx context.Context, f *model.Follow) error
	GetUserIdsByFolloweredId(ctx context.Context, userId, limit int64, cursor time.Time) ([]*model.FolloweringIdWithTime, int64, error)
	GetUserIdsByFolloweringId(ctx context.Context, userId, limit int64, cursor time.Time) ([]*model.FolloweredIdWithTime, int64, error)
}

type RpcPort interface {
	GetUsersByIdsRPC(ctx context.Context, ids []int64) ([]*model.UserProfile, error)
}

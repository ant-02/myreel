package repository

import (
	"context"
	"myreel/app/follow/domain/model"
)

type FollowDB interface {
	Magrate() error
	GetFollowByUserIdAndToUserId(ctx context.Context, userId, toUserId int64) (*model.Follow, error)
	SetFollowStatus(ctx context.Context, id, status int64) error
	CreateFollow(ctx context.Context, f *model.Follow) error
}

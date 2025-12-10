package repository

import (
	"context"
	"myreel/app/user/domain/model"
)

type UserDB interface {
	Magrate() error
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetUserById(ctx context.Context, id int64) (*model.User, error)
	SetAvatar(ctx context.Context, id int64, url string) error
	GetUserIdByUserName(ctx context.Context, username string) (int64, error)
	GetUsersByIds(ctx context.Context, ids []int64) ([]*model.UserProfile, error)
}

type UserCache interface {
	IsExist(ctx context.Context, key string) bool
	SetUserBaned(ctx context.Context, uid int64) error
	DeleteUserBaned(ctx context.Context, uid int64) error
	GetBannedUserIDs(ctx context.Context, pattern string) ([]int64, error)
	UserBanedKey(uid int64) string
	UserLoginKey(uid int64) string
	SetUserLogin(ctx context.Context, key string, token string) error
	GetUserLogin(ctx context.Context, key string) (string, error)
	DeleteUserLogin(ctx context.Context, key string) error
}

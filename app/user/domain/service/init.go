package service

import (
	"context"
	"myreel/app/user/domain/model"
	"myreel/app/user/domain/repository"
	"myreel/pkg/util"
)

type userService struct {
	db    repository.UserDB
	cache repository.UserCache
	sf    *util.Snowflake
}

type UserService interface {
	GenerateUserId() (int64, error)
	EncryptPassword(pwd string) (string, error)
	UserRegister(ctx context.Context, username, password string) (*model.User, error)
	IsBaned(ctx context.Context, uid int64) bool
	CheckPassword(ctx context.Context, ePwd, pwd string) error
	UserLogin(ctx context.Context, uid int64) (*model.Token, error)
	GetUserById(ctx context.Context, uid int64) (*model.User, error)
}

func NewUserService(db repository.UserDB, sf *util.Snowflake, cache repository.UserCache) UserService {
	if db == nil {
		panic("userService`s db should not be nil")
	}

	svc := &userService{
		db:    db,
		sf:    sf,
		cache: cache,
	}
	return svc
}

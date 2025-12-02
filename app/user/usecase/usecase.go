package usecase

import (
	"context"
	"myreel/app/user/domain/model"
	"myreel/app/user/domain/repository"
	"myreel/app/user/domain/service"
	"myreel/pkg/upyun"
)

type useCase struct {
	db    repository.UserDB
	svc   service.UserService
	cache repository.UserCache
}

type UserUseCase interface {
	Register(ctx context.Context, username, password string) error
	Login(ctx context.Context, username, password string) (*model.User, *model.Token, error)
	GetUserById(ctx context.Context, uid int64) (*model.User, error)
	Refresh(ctx context.Context, token string, uid int64) (string, error)
	GetLoadToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error)
}

func NewUserUseCase(db repository.UserDB, svc service.UserService, cache repository.UserCache) *useCase {
	return &useCase{
		db:    db,
		svc:   svc,
		cache: cache,
	}
}

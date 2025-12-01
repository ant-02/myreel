package usecase

import (
	"context"
	"myreel/app/user/domain/model"
	"myreel/app/user/domain/repository"
	"myreel/app/user/domain/service"
)

type useCase struct {
	db    repository.UserDB
	svc   service.UserService
	cache repository.UserCache
}

type UserUseCase interface {
	Register(ctx context.Context, username, password string) error
	Login(ctx context.Context, username, password string) (*model.User, *model.Token, error)
}

func NewUserUseCase(db repository.UserDB, svc service.UserService, cache repository.UserCache) *useCase {
	return &useCase{
		db:    db,
		svc:   svc,
		cache: cache,
	}
}

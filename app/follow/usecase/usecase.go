package usecase

import (
	"myreel/app/follow/domain/repository"
	"myreel/app/follow/domain/service"
)

type useCase struct {
	db  repository.FollowDB
	svc service.FollowService
}

type FollowUseCase interface {
}

func NewFollowUseCase(db repository.FollowDB, svc service.FollowService) FollowUseCase {
	return &useCase{
		db:  db,
		svc: svc,
	}
}

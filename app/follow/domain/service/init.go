package service

import (
	"myreel/app/follow/domain/repository"
	"myreel/pkg/util"
)

type FollowService interface {
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

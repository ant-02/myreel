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
	GetUsersByFolloweredId(ctx context.Context, userId, cursor, limit int64) ([]*model.UserProfile, *model.Pagination, error)
	GetUsersByFolloweringId(ctx context.Context, userId, cursor, limit int64) ([]*model.UserProfile, *model.Pagination, error)
	GetFriendsById(ctx context.Context, id, cursor, limit int64) ([]*model.UserProfile, *model.Pagination, error)
	GenerateGroupId() (int64, error)
	CreateGroup(ctx context.Context, g *model.Group) error
	GenerateGroupMemberId() (int64, error)
	CreateGroupMember(ctx context.Context, gm *model.GroupMember) error
	GetGroupByJoined(ctx context.Context, userId int64) ([]*model.Group, error)
	GetGroupByCreator(ctx context.Context, creatorId int64) ([]*model.Group, error)
}

type followService struct {
	db  repository.FollowDB
	sf  *util.Snowflake
	rpc repository.RpcPort
}

func NewFollowService(db repository.FollowDB, sf *util.Snowflake, rpc repository.RpcPort) FollowService {
	if db == nil {
		panic("followService`s db should not be nil")
	}

	svc := &followService{
		db:  db,
		sf:  sf,
		rpc: rpc,
	}
	return svc
}

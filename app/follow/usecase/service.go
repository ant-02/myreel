package usecase

import (
	"context"
	"errors"
	"myreel/app/follow/domain/model"
	"myreel/pkg/errno"
)

func (uc *useCase) FollowAction(ctx context.Context, userId, toUserId, actionType int64) error {
	err := uc.svc.FollowAction(ctx, userId, toUserId, actionType)
	if err != nil && errors.Is(err, errno.FollowNotFound) {
		id, err := uc.svc.GenerateFollowId()
		if err != nil {
			return err
		}
		return uc.svc.CreateFollow(ctx, &model.Follow{
			Id:            id,
			FolloweringId: userId,
			FolloweredId:  toUserId,
			Status:        actionType,
		})
	}
	return err
}

func (uc *useCase) GetUsersByFolloweredId(ctx context.Context, userId, cursor, limit int64) ([]*model.UserProfile, *model.Pagination, error) {
	return uc.svc.GetUsersByFolloweredId(ctx, userId, cursor, limit)
}

func (uc *useCase) GetUsersByFolloweringId(ctx context.Context, userId, cursor, limit int64) ([]*model.UserProfile, *model.Pagination, error) {
	return uc.svc.GetUsersByFolloweringId(ctx, userId, cursor, limit)
}

func (uc *useCase) GetFriendsById(ctx context.Context, id, cursor, limit int64) ([]*model.UserProfile, *model.Pagination, error) {
	return uc.svc.GetFriendsById(ctx, id, cursor, limit)
}

func (uc *useCase) CreateGroup(ctx context.Context, creatorId int64, name string, userIds ...int64) error {
	groupId, err := uc.svc.GenerateGroupId()
	if err != nil {
		return err
	}

	err = uc.svc.CreateGroup(ctx, &model.Group{
		Id:        groupId,
		Name:      name,
		CreatorId: creatorId,
	})
	if err != nil {
		return err
	}

	gmId, err := uc.svc.GenerateGroupMemberId()
	if err != nil {
		return err
	}
	err = uc.svc.CreateGroupMember(ctx, &model.GroupMember{
		Id:      gmId,
		GroupId: groupId,
		UserId:  creatorId,
	})
	if err != nil {
		return err
	}

	l := len(userIds)
	if l > 0 {
		for _, v := range userIds {
			gmId, err := uc.svc.GenerateGroupMemberId()
			if err != nil {
				return err
			}
			err = uc.svc.CreateGroupMember(ctx, &model.GroupMember{
				Id:      gmId,
				GroupId: groupId,
				UserId:  v,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (uc *useCase) GetGroupByJoined(ctx context.Context, userId int64) ([]*model.Group, error) {
	return uc.svc.GetGroupByJoined(ctx, userId)
}

func (uc *useCase) GetGroupByCreator(ctx context.Context, creatorId int64) ([]*model.Group, error) {
	return uc.svc.GetGroupByCreator(ctx, creatorId)
}

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

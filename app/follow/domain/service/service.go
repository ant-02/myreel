package service

import (
	"context"
	"myreel/app/follow/domain/model"
	"myreel/pkg/errno"
)

func (fs *followService) FollowAction(ctx context.Context, userId, toUserId, actionType int64) error {
	follow, err := fs.db.GetFollowByUserIdAndToUserId(ctx, userId, toUserId)
	if err != nil {
		return err
	}
	if actionType == follow.Status {
		return nil
	}
	err = fs.db.SetFollowStatus(ctx, follow.Id, actionType)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to set follow status")
	}
	return nil
}

func (fs *followService) GenerateFollowId() (int64, error) {
	id, err := fs.sf.Generate()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (fs *followService) CreateFollow(ctx context.Context, f *model.Follow) error {
	if err := fs.db.CreateFollow(ctx, f); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to create follow").WithError(err)
	}
	return nil
}

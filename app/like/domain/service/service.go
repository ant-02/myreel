package service

import (
	"context"
	"myreel/app/like/domain/model"
	"myreel/pkg/errno"
)

func (us *likeService) GenerateLikeId() (int64, error) {
	id, err := us.sf.Generate()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (ls *likeService) CreateLike(ctx context.Context, l *model.Like) error {
	if err := ls.db.CreateLike(l); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to create like").WithError(err)
	}
	return nil
}

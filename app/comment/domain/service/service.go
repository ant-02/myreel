package service

import (
	"context"
	"myreel/app/comment/domain/model"
	"myreel/pkg/errno"
)

func (cs *commentService) GenerateLikeId() (int64, error) {
	id, err := cs.sf.Generate()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (cs *commentService) CommentPublish(ctx context.Context, comment *model.Comment) error {
	if err := cs.db.CreateComment(ctx, comment); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to create comment").WithError(err)
	}
	return nil
}

func (cs *commentService) AddChildCount(ctx context.Context, commentId int64) error {
	if err := cs.db.AddChildCount(ctx, commentId); err != nil {
		errno.NewErrNo(errno.InternalServiceErrorCode, "failed to add comment child count").WithError(err)
	}
	return nil
}

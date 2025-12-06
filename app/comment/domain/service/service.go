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

func (cs *commentService) GetCommentListByVideoId(ctx context.Context, videoId, cursor, limit int64) ([]*model.Comment, *model.Pagination, error) {
	videos, total, err := cs.db.GetCommentListByVideoId(ctx, videoId, cursor, limit)
	if err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get comments by video id").WithError(err)
	}
	l := len(videos)
	var nextCursor int64
	if l > 0 {
		nextCursor = videos[l-1].CreatedAt
	} else {
		nextCursor = cursor
	}
	return videos, &model.Pagination{
		NextCursor: nextCursor,
		PrevCursor: cursor,
		Total:      total,
	}, nil
}

func (cs *commentService) GetCommentListByCommentId(ctx context.Context, commentId, cursor, limit int64) ([]*model.Comment, *model.Pagination, error) {
	videos, total, err := cs.db.GetCommentListByCommentId(ctx, commentId, cursor, limit)
	if err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get comments by comment id").WithError(err)
	}
	l := len(videos)
	var nextCursor int64
	if l > 0 {
		nextCursor = videos[l-1].CreatedAt
	} else {
		nextCursor = cursor
	}
	return videos, &model.Pagination{
		NextCursor: nextCursor,
		PrevCursor: cursor,
		Total:      total,
	}, nil
}

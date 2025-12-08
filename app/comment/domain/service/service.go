package service

import (
	"context"
	"errors"
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

func (cs *commentService) DeleteCommentById(ctx context.Context, id, uid int64) error {
	comment, err := cs.db.GetCommentById(ctx, id)
	if err != nil {
		if errors.Is(err, errno.CommentNotFound) {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "comment not found")
		}
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get comment parent id by id").WithError(err)
	}

	if comment.Uid != uid {
		return errno.NewErrNo(errno.AuthInvalidCode, "you can't delete other's comment")
	}

	err = cs.db.DeleteCommentById(ctx, id)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to delete comment by id").WithError(err)
	}

	err = cs.db.SubtractChildCount(ctx, comment.ParentId)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to subtract parent comment's child count").WithError(err)
	}

	return nil
}

func (cs *commentService) DeleteCommentsByVideoId(ctx context.Context, videoId, uid int64) error {
	if err := cs.vRpc.CheckVideoUser(ctx, videoId, uid); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "can not verify video's user").WithError(err)
	}

	if err := cs.db.DeleteCommentsByVideoId(ctx, videoId); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to delete comments by video id").WithError(err)
	}
	return nil
}

func (cs *commentService) AddLikeCount(ctx context.Context, id int64) error {
	if err := cs.db.AddLikeCount(ctx, id); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to add comment's like count").WithError(err)
	}
	return nil
}

func (cs *commentService) SubtractLikeCount(ctx context.Context, id int64) error {
	if err := cs.db.SubtractLikeCount(ctx, id); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to subtract comment's like count").WithError(err)
	}
	return nil
}

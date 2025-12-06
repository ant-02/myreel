package usecase

import (
	"context"
	"myreel/app/comment/domain/model"
	"myreel/pkg/errno"
)

func (uc *useCase) CommentPublish(ctx context.Context, videoId, commentId, userId int64, content string) error {
	id, err := uc.svc.GenerateLikeId()
	if err != nil {
		return err
	}

	if commentId != 0 {
		err = uc.svc.AddChildCount(ctx, commentId)
		if err != nil {
			return err
		}
	} else {
		err = uc.vRpc.AddCommentCount(ctx, videoId)
		if err != nil {
			return err
		}
	}

	return uc.svc.CommentPublish(ctx, &model.Comment{
		Id:       id,
		VideoId:  videoId,
		Uid:      userId,
		ParentId: commentId,
		Content:  content,
	})
}

func (uc *useCase) GetCommentList(ctx context.Context, videoId, commentId, cursor, limit int64) ([]*model.Comment, *model.Pagination, error) {
	if commentId == 0 {
		return uc.svc.GetCommentListByVideoId(ctx, videoId, cursor, limit)
	}
	return uc.svc.GetCommentListByCommentId(ctx, commentId, cursor, limit)
}

func (uc *useCase) DeleteComment(ctx context.Context, videoId, commentId, uid int64) error {
	if commentId == 0 && videoId == 0 {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "video id and comment id all empty")
	}
	if commentId == 0 {
		return uc.svc.DeleteCommentsByVideoId(ctx, videoId, uid)
	} else {
		return uc.svc.DeleteCommentById(ctx, commentId, uid)
	}
}

package usecase

import (
	"context"
	"myreel/app/like/domain/model"
	"myreel/pkg/errno"
)

func (uc *useCase) LikeAction(ctx context.Context, videoId, commentId, uid, actionType int64) error {
	if videoId != 0 {
		l, err := uc.db.GetVideoLike(ctx, videoId, uid)
		if err == nil {
			if l.Status == actionType {
				return nil
			}
			err = uc.db.SetLikeStatus(ctx, l.Id, actionType)
			if err != nil {
				return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to set like status").WithError(err)
			}
		}

		lid, err := uc.svc.GenerateLikeId()
		if err != nil {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to generate like id").WithError(err)
		}

		err = uc.svc.CreateLike(ctx, &model.Like{
			Id:      lid,
			Uid:     uid,
			VideoId: videoId,
			Status:  actionType,
		})
		return err
	} else if commentId != 0 {
		l, err := uc.db.GetCommentLike(ctx, commentId, uid)
		if err == nil {
			if l.Status == actionType {
				return nil
			}
			err = uc.db.SetLikeStatus(ctx, l.Id, actionType)
			if err != nil {
				return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to set like status").WithError(err)
			}
		}

		lid, err := uc.svc.GenerateLikeId()
		if err != nil {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to generate like id").WithError(err)
		}

		err = uc.svc.CreateLike(ctx, &model.Like{
			Id:        lid,
			Uid:       uid,
			CommentId: commentId,
			Status:    actionType,
		})
		return err
	}
	return errno.NewErrNo(errno.InternalServiceErrorCode, "video id and comment id all is empty")
}

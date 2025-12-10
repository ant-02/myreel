package usecase

import (
	"context"
	"errors"
	"fmt"
	"myreel/app/like/domain/model"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
)

func (uc *useCase) LikeAction(ctx context.Context, videoId, commentId, uid, actionType int64) error {
	if (videoId == 0 && commentId == 0) || (videoId != 0 && commentId != 0) {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "video id and comment id all is empty")
	}
	var err error
	if videoId != 0 {
		err = uc.svc.VideoUserLikeAction(ctx, videoId, uid, actionType)
		if err == nil || !errors.Is(err, errno.LikeNotFound) {
			return err
		}

		err = uc.lRpc.VideoLikeAction(ctx, videoId, actionType)
		if err != nil {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to action video like").WithError(err)
		}

		var lid int64
		lid, err = uc.svc.GenerateLikeId()
		if err != nil {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to generate like id").WithError(err)
		}

		err = uc.svc.CreateLike(ctx, &model.Like{
			Id:      lid,
			Uid:     uid,
			VideoId: videoId,
			Status:  actionType,
		})
	} else {
		err = uc.svc.CommentUserLikeAction(ctx, commentId, uid, actionType)
		if err == nil || !errors.Is(err, errno.LikeNotFound) {
			return err
		}

		err = uc.lRpc.CommentLikeAction(ctx, commentId, actionType)
		if err != nil {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to action comment like").WithError(err)
		}
		var lid int64
		lid, err = uc.svc.GenerateLikeId()
		if err != nil {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to generate like id").WithError(err)
		}

		err = uc.svc.CreateLike(ctx, &model.Like{
			Id:        lid,
			Uid:       uid,
			CommentId: commentId,
			Status:    actionType,
		})
	}
	return err
}

func (uc *useCase) GetVideosByUserLike(ctx context.Context, userId, cursor, limit int64) ([]*model.Video, *model.Pagination, error) {
	key := fmt.Sprintf("%s%d", constants.RedisUserLikeKey, userId)

	ids, err := uc.cache.GetVideoIdFromUserLike(ctx, key, cursor, limit)
	if err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get video id list by user like").WithError(err)
	}

	total, err := uc.cache.GetVideoLikeCount(ctx, key)
	if err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get user video count").WithError(err)
	}

	videos, err := uc.lRpc.GetVideosByIds(ctx, ids)
	if err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get video list by id list").WithError(err)
	}

	var nextCursor int64
	l := len(videos)
	if l == 0 {
		nextCursor = cursor
	} else {
		nextCursor = videos[l-1].CreatedAt
	}

	return videos, &model.Pagination{
		NextCursor: nextCursor,
		PrevCursor: cursor,
		Total:      total,
	}, nil
}

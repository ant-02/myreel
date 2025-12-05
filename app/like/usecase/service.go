package usecase

import (
	"context"
	"fmt"
	"myreel/app/like/domain/model"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
	"time"
)

func (uc *useCase) VideoUserLikeAction(ctx context.Context, videoId, uid, actionType int64) error {
	key := fmt.Sprintf("%s%d", constants.RedisUserLikeKey, uid)
	exist, err := uc.cache.IsExist(ctx, key, videoId)
	if err != nil {
		return err
	}
	if actionType == 0 {
		if !exist {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "user already cancel like this video")
		}
		err = uc.cache.RemVideoLikeFromUser(ctx, key, videoId)
		if err != nil {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to remove video like from user set").WithError(err)
		}
	} else {
		if exist {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "user already like this video")
		}
		err = uc.cache.AddVideoLikeToUser(ctx, key, float64(time.Now().Unix()), videoId)
		if err != nil {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to add video like from user set").WithError(err)
		}
	}

	err = uc.lRpc.VideoLikeAction(ctx, videoId, actionType)
	if err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to action video like").WithError(err)
	}

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
}

func (uc *useCase) CommentUserLikeAction(ctx context.Context, commentId, uid, actionType int64) error {
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

func (uc *useCase) LikeAction(ctx context.Context, videoId, commentId, uid, actionType int64) error {
	if (videoId == 0 && commentId == 0) || (videoId != 0 && commentId != 0) {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "video id and comment id all is empty")
	}
	if videoId != 0 {
		return uc.VideoUserLikeAction(ctx, videoId, uid, actionType)
	} else {
		return uc.CommentUserLikeAction(ctx, commentId, uid, actionType)
	}

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

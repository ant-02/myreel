package service

import (
	"context"
	"fmt"
	"myreel/app/like/domain/model"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
	"strconv"
	"time"
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

func (ls *likeService) VideoUserLikeAction(ctx context.Context, videoId, uid, actionType int64) error {
	key := fmt.Sprintf("%s%d", constants.RedisUserLikeKey, uid)
	exist, err := ls.cache.IsExist(ctx, key, strconv.FormatInt(videoId, 10))
	if err != nil {
		return err
	}
	if actionType == 0 {
		if !exist {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "user already cancel like this video")
		}
		err = ls.cache.RemVideoLikeFromUser(ctx, key, videoId)
		if err != nil {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to remove video like from user set").WithError(err)
		}
	} else {
		if exist {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "user already like this video")
		}
		err = ls.cache.AddVideoLikeToUser(ctx, key, float64(time.Now().Unix()), videoId)
		if err != nil {
			return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to add video like from user set").WithError(err)
		}
	}

	l, err := ls.db.GetVideoLike(ctx, videoId, uid)
	if err == nil {
		if l.Status != actionType {
			err = ls.db.SetLikeStatus(ctx, l.Id, actionType)
			if err != nil {
				return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to set like status").WithError(err)
			}
			err = ls.lRpc.VideoLikeAction(ctx, videoId, actionType)
			if err != nil {
				return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to action video like").WithError(err)
			}
		}
		return nil
	}

	return err
}

func (ls *likeService) CommentUserLikeAction(ctx context.Context, commentId, uid, actionType int64) error {
	l, err := ls.db.GetCommentLike(ctx, commentId, uid)
	if err == nil {
		if l.Status != actionType {
			err = ls.db.SetLikeStatus(ctx, l.Id, actionType)
			if err != nil {
				return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to set like status").WithError(err)
			}
			err := ls.lRpc.CommentLikeAction(ctx, commentId, actionType)
			if err != nil {
				return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to action comment like").WithError(err)
			}
		}
		return nil
	}

	return err
}

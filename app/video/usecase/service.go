package usecase

import (
	"context"
	"fmt"
	"myreel/app/video/domain/model"
	"myreel/config"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
	"myreel/pkg/upyun"
)

func (us *useCase) GetVideosByLatestTime(ctx context.Context, latestTime string) ([]*model.Video, error) {
	return us.svc.GetVideosByLatestTime(ctx, latestTime)
}

func (us *useCase) GetVideoUplaodToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error) {
	return us.svc.GetVideoUploadToken(ctx, suffix, uid)
}

func (us *useCase) GetVideoCoverUplaodToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error) {
	return us.svc.GetVideoCoverUploadToken(ctx, suffix, uid)
}

func (us *useCase) SaveVideo(ctx context.Context, video *model.Video) error {
	var err error
	video.Id, err = us.svc.GenerateVideoId()
	if err != nil {
		return err
	}

	video.VideoUrl = fmt.Sprintf("%s%s", config.Upyun.Domain, video.VideoUrl)
	video.CoverUrl = fmt.Sprintf("%s%s", config.Upyun.Domain, video.CoverUrl)
	err = us.svc.SaveVideo(ctx, video)
	if err != nil {
		return err
	}

	return nil
}

func (us *useCase) GetVideosByUserId(ctx context.Context, uid, cursor, limit int64) ([]*model.Video, *model.Pagination, error) {
	return us.svc.GetVideosByUserId(ctx, uid, cursor, limit)
}

func (us *useCase) GetVideosGroupByVisitCount(ctx context.Context, cursor, limit int64) ([]*model.Video, *model.Pagination, error) {
	ids, err := us.cache.GetPopularVideos(ctx, constants.RedisVideoPopularKey, cursor, limit)
	if err != nil {
		return nil, nil, err
	}

	videos, err := us.svc.GetVideosByIds(ctx, ids)
	if err != nil {
		return nil, nil, err
	}

	l := len(videos)
	var nextCursor int64
	if l > 0 {
		nextCursor = us.svc.CalculateHotScore(videos[l-1])
	} else {
		nextCursor = cursor
	}

	return videos, &model.Pagination{
		NextCursor: nextCursor,
		PrevCursor: cursor,
		Total:      constants.RedisVideoPopCount,
	}, nil
}

func (us *useCase) GetVideosByKeywords(ctx context.Context, keywords, username string, fromDate, toDate, cursor, limit int64) ([]*model.Video, *model.Pagination, error) {
	var id int64 = 0
	var err error
	if username != "" {
		id, err = us.vRpc.GetUserIdByUsername(ctx, username)
		if err != nil {
			return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get user id by username").WithError(err)
		}
	}
	return us.svc.GetVideosByKeywords(ctx, keywords, fromDate, toDate, cursor, id, limit)
}

func (us *useCase) VideoLikeAction(ctx context.Context, videoId, actionType int64) error {
	if actionType == 0 {
		return us.svc.DecrVideoLike(ctx, videoId)
	} else {
		return us.svc.IncrVideoLike(ctx, videoId)
	}
}

func (us *useCase) GetVideosByIds(ctx context.Context, ids []int64) ([]*model.Video, error) {
	return us.svc.GetVideosByIds(ctx, ids)
}

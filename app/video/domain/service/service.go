package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"myreel/app/video/domain/model"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
	"myreel/pkg/upyun"
	"strconv"
	"time"
)

func (vs *videoService) GetVideosByLatestTime(ctx context.Context, latestTime string) ([]*model.Video, error) {
	t, err := strconv.ParseInt(latestTime, 10, 64)
	if err != nil {
		return nil, errno.Errorf(errno.InternalServiceErrorCode, "failed to parse time: %v", err)
	}

	lt := time.Unix(t, 0)
	videos, err := vs.db.GetVideosByLatestTime(ctx, lt)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get videos by latest time").WithError(err)
	}

	return videos, nil
}

func (vs *videoService) GetVideoUploadToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error) {
	saveKey := fmt.Sprintf("%s/%s/%d%s", constants.UpyunVideoPath, time.Now().Format("2006/01/02"), uid, suffix)
	up, err := upyun.GeneratePolicyAndSignature(uid, saveKey, "", nil)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get upyun token").WithError(err)
	}

	return up, nil
}

func (vs *videoService) GetVideoCoverUploadToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error) {
	saveKey := fmt.Sprintf("%s/%s/%d%s", constants.UpyunVideoCoverPath, time.Now().Format("2006/01/02"), uid, suffix)
	up, err := upyun.GeneratePolicyAndSignature(uid, saveKey, "", nil)
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get upyun token").WithError(err)
	}

	return up, nil
}

func (us *videoService) GenerateVideoId() (int64, error) {
	id, err := us.sf.Generate()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (vs *videoService) SaveVideo(ctx context.Context, video *model.Video) error {
	if err := vs.db.CreateVideo(ctx, video); err != nil {
		return errno.NewErrNo(errno.InternalServiceErrorCode, "failed to save video").WithError(err)
	}
	return nil
}

func (vs *videoService) GetVideosByUserId(ctx context.Context, uid, cursor, limit int64) ([]*model.Video, *model.Pagination, error) {
	videos, total, err := vs.db.GetVideosByUid(ctx, uid, cursor, limit)
	if err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get videos info by user id").WithError(err)
	}

	pagination := &model.Pagination{
		PrevCursor: cursor,
		Total:      total,
	}
	l := len(videos)
	if l == 0 {
		pagination.NextCursor = cursor
	} else {
		pagination.NextCursor = videos[l-1].Id
	}

	return videos, pagination, nil
}

func (vs *videoService) GetVideosByIds(ctx context.Context, ids []int64) ([]*model.Video, error) {
	if ids == nil {
		return nil, nil
	}

	l := len(ids)
	videos := make([]*model.Video, l)
	for i, id := range ids {
		key := fmt.Sprintf("%s/%d", constants.RedisVideoKey, id)
		if exist := vs.cache.IsExist(ctx, key); exist {
			video, err := vs.cache.GetVideo(ctx, key)
			if err != nil {
				return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get video").WithError(err)
			}
			videos[i] = video
		} else {
			video, err := vs.db.GetVideoById(ctx, id)
			if err != nil {
				return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get video").WithError(err)
			}
			vj, err := json.Marshal(video)
			if err != nil {
				return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to marshal video").WithError(err)
			}
			err = vs.cache.AddVideoWithTLL(ctx, key, vj, constants.RedisVideoExpireTime)
			if err != nil {
				return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to add video to redis").WithError(err)
			}
			videos[i] = video
		}
	}
	return videos, nil
}

func (vs *videoService) GetVideosByKeywords(ctx context.Context, keywords string, fromDate, toDate, cursor, uid, limit int64) ([]*model.Video, *model.Pagination, error) {
	var ft, tt, ct time.Time
	ft = time.Unix(fromDate, 0)
	if toDate == 0 {
		tt = time.Now()
	} else {
		tt = time.Unix(toDate, 0)
	}
	if cursor == 0 {
		ct = tt
	} else {
		ct = time.Unix(cursor, 0)
	}

	videos, total, err := vs.db.GetVideosByKeywords(ctx, keywords, ft, tt, ct, uid, limit)
	if err != nil {
		return nil, nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get video by keywords").WithError(err)
	}
	pagination := &model.Pagination{
		PrevCursor: cursor,
		Total:      total,
	}
	l := len(videos)
	if l == 0 {
		pagination.NextCursor = cursor
	} else {
		pagination.NextCursor = videos[l-1].CreatedAt
	}

	return videos, pagination, nil
}

func (vs *videoService) CalculateHotScore(video *model.Video) int64 {	
	if video == nil {
		return 0
	}

	baseScore := video.LikeCount*constants.VideoLikeGravity + video.CommentCount*constants.VideoCommentGarvity + video.VisitCount*constants.VideoVisitGarvity
	if baseScore <= 0 {
		return 0
	}

	now := time.Now().Unix()
	ageSeconds := now - video.CreatedAt
	if ageSeconds < 0 {
		ageSeconds = 0
	}

	ageHours := float64(ageSeconds) / 3600.0

	denominator := math.Pow(ageHours+2, constants.VideoCreatedAtGarvity)
	floatScore := float64(baseScore) / denominator
	scaled := floatScore * 1000.0
	return int64(math.Round(scaled))
}

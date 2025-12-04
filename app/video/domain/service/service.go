package service

import (
	"context"
	"fmt"
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

	videos, err := vs.db.GetVideosByLatestTime(ctx, t)
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

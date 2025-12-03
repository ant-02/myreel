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
	up, err := upyun.GeneratePolicyAndSignature(uid, saveKey, constants.UpyunUserAvatarNotifyPath, strconv.FormatInt(uid, 10))
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalServiceErrorCode, "failed to get upyun token").WithError(err)
	}

	return up, nil
}

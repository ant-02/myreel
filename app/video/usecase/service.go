package usecase

import (
	"context"
	"myreel/app/video/domain/model"
	"myreel/pkg/upyun"
)

func (us *useCase) GetVideosByLatestTime(ctx context.Context, latestTime string) ([]*model.Video, error) {
	return us.svc.GetVideosByLatestTime(ctx, latestTime)
}

func (us *useCase) GetVideoUplaodToken(ctx context.Context, suffix string, uid int64) (*upyun.UpyunToken, error) {
	return us.svc.GetVideoUploadToken(ctx, suffix, uid)
}

package usecase

import (
	"context"
	"fmt"
	"myreel/app/video/domain/model"
	"myreel/config"
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

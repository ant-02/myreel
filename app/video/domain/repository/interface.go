package repository

import (
	"context"
	"myreel/app/video/domain/model"
)

type VideoDB interface {
	Magrate() error
	GetVideosByLatestTime(ctx context.Context, latestTime int64) ([]*model.Video, error)
	CreateVideo(ctx context.Context, video *model.Video) error
}

type VideoCache interface {
}

package repository

import (
	"context"
	"myreel/app/video/domain/model"
)

type VideoDB interface {
	GetVideosByLatestTime(ctx context.Context, latestTime int64) ([]*model.Video, error)
}

type VideoCache interface {
}

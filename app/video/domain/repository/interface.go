package repository

import (
	"context"
	"myreel/app/video/domain/model"
)

type VideoDB interface {
	Magrate() error
	GetVideosByLatestTime(ctx context.Context, latestTime int64) ([]*model.Video, error)
	CreateVideo(ctx context.Context, video *model.Video) error
	GetVideosByUid(ctx context.Context, uid, cursor, limit int64) ([]*model.Video, int64, error)
	GetVideosGroupByVisitCount(ctx context.Context, cursor, limit int64) ([]*model.Video, int64, error)
}

type VideoCache interface {
}

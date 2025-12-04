package repository

import (
	"context"
	"myreel/app/video/domain/model"
	"time"
)

type VideoDB interface {
	Magrate() error
	GetVideosByLatestTime(ctx context.Context, latestTime time.Time) ([]*model.Video, error)
	CreateVideo(ctx context.Context, video *model.Video) error
	GetVideosByUid(ctx context.Context, uid, cursor, limit int64) ([]*model.Video, int64, error)
	GetVideosGroupByVisitCount(ctx context.Context, cursor, limit int64) ([]*model.Video, int64, error)
	GetVideosByKeywords(ctx context.Context, keywords string, fromDate, toDate, cursor time.Time, uid, limit int64) ([]*model.Video, int64, error)
}

type VideoCache interface {
}

type RpcPort interface {
	GetUserIdByUsername(ctx context.Context, username string) (int64, error)
}

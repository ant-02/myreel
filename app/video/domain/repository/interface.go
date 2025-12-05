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
	GetVideosByKeywords(ctx context.Context, keywords string, fromDate, toDate, cursor time.Time, uid, limit int64) ([]*model.Video, int64, error)
	GetVideoById(ctx context.Context, id int64) (*model.Video, error)
	AddLikeCount(ctx context.Context, id int64) error
	SubtractLikeCount(ctx context.Context, id int64) error
}

type VideoCache interface {
	IsExist(ctx context.Context, key string) bool
	AddPopularVideoId(ctx context.Context, key string, score float64, member interface{}) error
	GetPopularVideos(ctx context.Context, key string, cursor, limit int64) ([]int64, error)
	CleanPopularVideos(ctx context.Context, key string, limit int64) error
	GetVideo(ctx context.Context, key string) (*model.Video, error)
	AddVideoWithTLL(ctx context.Context, key string, val interface{}, ttl time.Duration) error
	VideoLikeDecr(ctx context.Context, key string) error
	VideoLikeIncr(ctx context.Context, key string) error
	AddVideoLike(ctx context.Context, key string, val interface{}, ttl time.Duration) error
}
type RpcPort interface {
	GetUserIdByUsername(ctx context.Context, username string) (int64, error)
}

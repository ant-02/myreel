package repository

import (
	"context"
	"myreel/app/like/domain/model"
)

type LikeDB interface {
	Magrate() error
	GetVideoLike(ctx context.Context, videoId, uid int64) (*model.Like, error)
	GetCommentLike(ctx context.Context, commentId, uid int64) (*model.Like, error)
	SetLikeStatus(ctx context.Context, id int64, status int64) error
	CreateLike(l *model.Like) error
}

type LikeCache interface {
	IsExist(ctx context.Context, key string, val interface{}) (bool, error)
	RemVideoLikeFromUser(ctx context.Context, key string, member interface{}) error
	AddVideoLikeToUser(ctx context.Context, key string, score float64, member interface{}) error
	GetVideoIdFromUserLike(ctx context.Context, key string, cursor, limit int64) ([]int64, error)
	GetVideoLikeCount(ctx context.Context, key string) (int64, error)
}

type RpcPort interface {
	VideoLikeAction(ctx context.Context, videoId, actionType int64) error
	GetVideosByIds(ctx context.Context, ids []int64) ([]*model.Video, error)
}

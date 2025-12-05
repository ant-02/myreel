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

type RpcPort interface {
}

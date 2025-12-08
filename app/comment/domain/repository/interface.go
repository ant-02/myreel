package repository

import (
	"context"
	"myreel/app/comment/domain/model"
)

type CommentDB interface {
	Magrate() error
	CreateComment(ctx context.Context, comment *model.Comment) error
	AddChildCount(ctx context.Context, commentId int64) error
	SubtractChildCount(ctx context.Context, commentId int64) error
	GetCommentListByVideoId(ctx context.Context, videoId, limit, cursor int64) ([]*model.Comment, int64, error)
	GetCommentListByCommentId(ctx context.Context, commentId, limit, cursor int64) ([]*model.Comment, int64, error)
	DeleteCommentById(ctx context.Context, id int64) error
	DeleteCommentsByVideoId(ctx context.Context, videoId int64) error
	GetCommentById(ctx context.Context, id int64) (*model.Comment, error)
	AddLikeCount(ctx context.Context, id int64) error
	SubtractLikeCount(ctx context.Context, id int64) error
}

type RpcPort interface {
	AddCommentCount(ctx context.Context, id int64) error
	CheckVideoUser(ctx context.Context, videoId, uid int64) error
}

package service

import (
	"context"
	"myreel/app/comment/domain/model"
	"myreel/app/comment/domain/repository"
	"myreel/pkg/util"
)

type commentService struct {
	db   repository.CommentDB
	sf   *util.Snowflake
	vRpc repository.RpcPort
}

type CommentService interface {
	GenerateLikeId() (int64, error)
	CommentPublish(ctx context.Context, comment *model.Comment) error
	AddChildCount(ctx context.Context, commentId int64) error
	GetCommentListByVideoId(ctx context.Context, videoId, cursor, limit int64) ([]*model.Comment, *model.Pagination, error)
	GetCommentListByCommentId(ctx context.Context, commentId, cursor, limit int64) ([]*model.Comment, *model.Pagination, error)
}

func NewCommentService(db repository.CommentDB, sf *util.Snowflake, vRpc repository.RpcPort) CommentService {
	if db == nil {
		panic("commentService`s db should not be nil")
	}

	svc := &commentService{
		db:   db,
		sf:   sf,
		vRpc: vRpc,
	}
	return svc
}

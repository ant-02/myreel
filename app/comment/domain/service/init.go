package service

import (
	"context"
	"myreel/app/comment/domain/model"
	"myreel/app/comment/domain/repository"
	"myreel/pkg/util"
)

type commentService struct {
	db repository.CommentDB
	sf *util.Snowflake
}

type CommentService interface {
	GenerateLikeId() (int64, error)
	CommentPublish(ctx context.Context, comment *model.Comment) error
	AddChildCount(ctx context.Context, commentId int64) error
}

func NewCommentService(db repository.CommentDB, sf *util.Snowflake) CommentService {
	if db == nil {
		panic("commentService`s db should not be nil")
	}

	svc := &commentService{
		db: db,
		sf: sf,
	}
	return svc
}

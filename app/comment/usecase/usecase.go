package usecase

import (
	"context"
	"myreel/app/comment/domain/repository"
	"myreel/app/comment/domain/service"
)

type useCase struct {
	db  repository.CommentDB
	svc service.CommentService
}

type CommentUseCase interface {
	CommentPublish(ctx context.Context, videoId, commentId, userId int64, content string) error
}

func NewCommentUseCase(db repository.CommentDB, svc service.CommentService) CommentUseCase {
	return &useCase{
		db:  db,
		svc: svc,
	}
}

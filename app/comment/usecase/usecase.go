package usecase

import (
	"context"
	"myreel/app/comment/domain/model"
	"myreel/app/comment/domain/repository"
	"myreel/app/comment/domain/service"
)

type useCase struct {
	db   repository.CommentDB
	svc  service.CommentService
	vRpc repository.RpcPort
}

type CommentUseCase interface {
	CommentPublish(ctx context.Context, videoId, commentId, userId int64, content string) error
	GetCommentList(ctx context.Context, videoId, commentId, cursor, limit int64) ([]*model.Comment, *model.Pagination, error)
	DeleteComment(ctx context.Context, videoId, commentId, uid int64) error
}

func NewCommentUseCase(db repository.CommentDB, svc service.CommentService, vRpc repository.RpcPort) CommentUseCase {
	return &useCase{
		db:   db,
		svc:  svc,
		vRpc: vRpc,
	}
}

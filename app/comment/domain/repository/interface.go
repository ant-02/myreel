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
}

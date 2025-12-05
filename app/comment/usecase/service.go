package usecase

import (
	"context"
	"myreel/app/comment/domain/model"
)

func (uc *useCase) CommentPublish(ctx context.Context, videoId, commentId, userId int64, content string) error {
	id, err := uc.svc.GenerateLikeId()
	if err != nil {
		return err
	}

	if commentId != 0 {
		err = uc.svc.AddChildCount(ctx, commentId); 
		if err != nil {
			return err
		}
	}

	return uc.svc.CommentPublish(ctx, &model.Comment{
		Id:       id,
		VideoId:  videoId,
		Uid:      userId,
		ParentId: commentId,
		Content:  content,
	})
}

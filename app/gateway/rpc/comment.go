package rpc

import (
	"context"
	api "myreel/app/gateway/model/api/comment"
	"myreel/kitex_gen/comment"
	"myreel/pkg/base/client"
	"myreel/pkg/errno"
	"myreel/pkg/logger"
	"myreel/pkg/util"
)

func InitCommentClient() {
	c, err := client.InitCommentRPC()
	if err != nil {
		logger.Fatalf("api.rpc.comment InitCommentRPC failed, err is %v", err)
	}
	commentClient = *c
}

func CommentPublishRPC(ctx context.Context, req *comment.CommentPublishRequest) error {
	resp, err := commentClient.CommentPublish(ctx, req)
	if err != nil {
		logger.Errorf("CommentPublishRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return nil
}

func GetCommentList(ctx context.Context, req *comment.CommentListRequest) (*api.CommentListResponse, error) {
	resp, err := commentClient.CommentList(ctx, req)
	if err != nil {
		logger.Errorf("GetCommentList: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	l := len(resp.Data.Items)
	comments := make([]*api.Comment, l)
	if l > 0 {
		for i, item := range resp.Data.Items {
			comments[i] = &api.Comment{
				Id:         item.Id,
				VideoId:    item.VideoId,
				Uid:        item.Uid,
				ParentId:   item.ParentId,
				LikeCount:  item.LikeCount,
				ChildCount: item.ChildCount,
				Content:    item.Content,
				CreatedAt:  item.CreatedAt,
				UpdatedAt:  item.UpdatedAt,
			}
		}
	}
	return &api.CommentListResponse{
		Items: comments,
		Pagination: &api.Pagination{
			NextCursor: resp.Data.Pagination.NextCursor,
			PrevCursor: resp.Data.Pagination.PrevCursor,
			Total:      resp.Data.Pagination.Total,
		},
	}, nil
}

func DeleteCommentRPC(ctx context.Context, req *comment.DeleteRequest) error {
	resp, err := commentClient.Delete(ctx, req)
	if err != nil {
		logger.Errorf("DeleteCommentRPC: RPC called failed: %v", err.Error())
		return  errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	return nil
}

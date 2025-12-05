package rpc

import (
	"context"
	"myreel/kitex_gen/comment"
	"myreel/pkg/base/client"
	"myreel/pkg/errno"
	"myreel/pkg/logger"
	"myreel/pkg/util"
	// api "myreel/app/gateway/model/api/comment"
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

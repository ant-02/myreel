package rpc

import (
	"context"
	"myreel/kitex_gen/follow"
	"myreel/pkg/base/client"
	"myreel/pkg/errno"
	"myreel/pkg/logger"
	"myreel/pkg/util"
)

func InitFollowClient() {
	c, err := client.InitFollowRPC()
	if err != nil {
		logger.Fatalf("api.rpc.like InitFollowRpc failed, err is %v", err)
	}
	followClient = *c
}

func FollowActionRPC(ctx context.Context, req *follow.FollowActionRequest) error {
	resp, err := followClient.FollowAction(ctx, req)
	if err != nil {
		logger.Errorf("FollowActionRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return nil
}

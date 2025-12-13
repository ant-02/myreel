package rpc

import (
	"context"
	"myreel/kitex_gen/chat"
	"myreel/pkg/base/client"
	"myreel/pkg/errno"
	"myreel/pkg/logger"
	"myreel/pkg/util"
)

func InitChatClient() {
	c, err := client.InitChatRPC()
	if err != nil {
		logger.Fatalf("api.rpc.chat InitChatRPC failed, err is %v", err)
	}
	chatClient = *c
}

func SendMessageRPC(ctx context.Context, req *chat.SendMessageRequest) error {
	resp, err := chatClient.SendMessage(ctx, req)
	if err != nil {
		logger.Errorf("SendMessageRPC: RPC called failed: %v", err.Error())
		return errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}
	return nil
}

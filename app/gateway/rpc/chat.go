package rpc

import (
	"context"
	api "myreel/app/gateway/model/api/chat"
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

func GetHistoryMessagesRPC(ctx context.Context, req *chat.GetHistoryRequest) (*api.MessageList, error) {
	resp, err := chatClient.GetHistory(ctx, req)
	if err != nil {
		logger.Errorf("GetHistoryMessagesRPC: RPC called failed: %v", err.Error())
		return nil, errno.InternalServiceError.WithError(err)
	}
	if !util.IsSuccess(resp.Base) {
		return nil, errno.InternalServiceError.WithMessage(resp.Base.Msg)
	}

	l := len(resp.Data.Messages)
	messages := make([]*api.Message, l)
	if l > 0 {
		for i, v := range resp.Data.Messages {
			messages[i] = &api.Message{
				Id:        v.Id,
				SenderId:  v.SenderId,
				TargetId:  v.TargetId,
				Content:   v.Content,
				CreatedAt: v.CreatedAt,
			}
		}
	}

	return &api.MessageList{
		Messages: messages,
		Pagination: &api.Pagination{
			NextCursor: resp.Data.Pagination.NextCursor,
			PrevCursor: resp.Data.Pagination.PrevCursor,
			Total:      resp.Data.Pagination.Total,
		},
	}, nil
}

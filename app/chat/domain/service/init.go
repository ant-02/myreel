package service

import (
	"context"
	"myreel/app/chat/domain/model"
	"myreel/app/chat/domain/repository"
	"myreel/pkg/util"
)

type chatService struct {
	db    repository.ChatDB
	cache repository.ChatCache
	sf    *util.Snowflake
	vRpc  repository.RpcPort
}

type ChatService interface {
	GenerateConversationID(chatType model.ChatType, id1, id2 int64) string
	GenerateMessageId() (int64, error)
	CreateMessage(ctx context.Context, msg *model.Message) error
	GetHistoryMessages(ctx context.Context, cursor, limit int64, conversationID string) ([]*model.Message, *model.Pagination, error)
	GetUnreadMessages(ctx context.Context, conversationID string) ([]*model.Message, *model.Pagination, error)
}

func NewChatService(db repository.ChatDB, sf *util.Snowflake, cache repository.ChatCache, vRpc repository.RpcPort) ChatService {
	if db == nil {
		panic("chatService`s db should not be nil")
	}

	svc := &chatService{
		db:    db,
		sf:    sf,
		cache: cache,
		vRpc:  vRpc,
	}
	return svc
}

package repository

import (
	"context"
	"myreel/app/chat/domain/model"
	"time"
)

type ChatDB interface {
	Magrate() error
	CreateMessage(ctx context.Context, msg *model.Message) error
	GetMessage(ctx context.Context, id int64) (*model.Message, error)
}

type ChatCache interface {
	IsExist(ctx context.Context, key string) bool
	HistoryKey(conversationID string) string
	UnreadKey(conversationID string) string
	MessageKey(id int64) string
	AddMessageId(ctx context.Context, key string, score float64, member int64) error
	AddMessageWithTTL(ctx context.Context, key string, message *model.Message, ttl time.Duration) error
	GetMessageIds(ctx context.Context, key string, cursor, limit int64) ([]int64, error)
	GetMessage(ctx context.Context, key string) (*model.Message, error)
	GetMessageCount(ctx context.Context, key string) (int64, error)
}

type RpcPort interface{}

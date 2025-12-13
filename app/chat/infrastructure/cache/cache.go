package cache

import (
	"context"
	"myreel/app/chat/domain/repository"

	"github.com/redis/go-redis/v9"
)

type chatCache struct {
	client *redis.Client
}

func NewChatCache(client *redis.Client) repository.ChatCache {
	return &chatCache{client: client}
}

func (c *chatCache) IsExist(ctx context.Context, key string) bool {
	return c.client.Exists(ctx, key).Val() == 1
}

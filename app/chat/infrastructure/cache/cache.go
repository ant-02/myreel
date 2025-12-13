package cache

import (
	"myreel/app/chat/domain/repository"

	"github.com/redis/go-redis/v9"
)

type chatCache struct {
	client *redis.Client
}

func NewChatCache(client *redis.Client) repository.ChatCache {
	return &chatCache{client: client}
}

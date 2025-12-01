package cache

import (
	"context"
	"myreel/app/user/domain/repository"

	"github.com/redis/go-redis/v9"
)

type userCache struct {
	client *redis.Client
}

func NewUserCache(client *redis.Client) repository.UserCache {
	return &userCache{client: client}
}

func (c *userCache) IsExist(ctx context.Context, key string) bool {
	return c.client.Exists(ctx, key).Val() == 1
}

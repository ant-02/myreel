package cache

import (
	"myreel/app/video/domain/repository"

	"github.com/redis/go-redis/v9"
)

type videoCache struct {
	client *redis.Client
}

func NewVideoCache(client *redis.Client) repository.VideoCache {
	return &videoCache{client: client}
}

package cache

import (
	"context"
	"myreel/app/chat/domain/model"
	"myreel/app/chat/infrastructure/cache/pack"
	"myreel/pkg/errno"
	"time"

	"github.com/redis/go-redis/v9"
)

func (cc *chatCache) AddMessageId(ctx context.Context, key string, score float64, member int64) error {
	if err := cc.client.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: member,
	}).Err(); err != nil {
		return errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to add message to zset: %v", err)
	}
	return nil
}

func (cc *chatCache) AddMessageWithTTL(ctx context.Context, key string, message *model.Message, ttl time.Duration) error {
	val, err := pack.MessageToMap(message)
	if err != nil {
		return errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to parse message to map: %v", err)
	}

	if err := cc.client.HSet(ctx, key, val).Err(); err != nil {
		return errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to set message to hash: %v", err)
	}

	err = cc.client.Expire(ctx, key, ttl).Err()
	if err != nil {
		return errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to add ttl: %v", err)
	}

	return nil
}

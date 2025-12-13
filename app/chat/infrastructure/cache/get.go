package cache

import (
	"context"
	"fmt"
	"myreel/app/chat/domain/model"
	"myreel/app/chat/infrastructure/cache/pack"
	"myreel/pkg/errno"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

func (cc *chatCache) GetMessageIds(ctx context.Context, key string, cursor, limit int64) ([]int64, error) {
	if cursor == 0 {
		cursor = time.Now().Unix()
	}
	ids, err := cc.client.ZRevRangeByScore(ctx, key, &redis.ZRangeBy{
		Max:   fmt.Sprintf("(%d", cursor),
		Count: limit,
	}).Result()
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis, failed to get message id: %v", err)
	}
	l := len(ids)
	result := make([]int64, l)
	if l > 0 {
		for i, v := range ids {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis, failed to get parse id: %v", err)
			}
			result[i] = id
		}
	}
	return result, nil
}

func (cc *chatCache) GetMessage(ctx context.Context, key string) (*model.Message, error) {
	val, err := cc.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to get message by id: %v", err)
	}
	
	message, err := pack.MapToMessage(val)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (cc *chatCache) GetMessageCount(ctx context.Context, key string) (int64, error) {
	return cc.client.ZCard(ctx, key).Result()
}

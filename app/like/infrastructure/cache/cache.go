package cache

import (
	"context"
	"errors"
	"fmt"
	"myreel/app/like/domain/repository"
	"myreel/pkg/errno"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type likeCache struct {
	client *redis.Client
}

func NewVideoCache(client *redis.Client) repository.LikeCache {
	return &likeCache{client: client}
}

func (c *likeCache) IsExist(ctx context.Context, key string, val string) (bool, error) {
	_, err := c.client.ZScore(ctx, key, val).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false, nil
		}
		return false, errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to check redis set member: %v", err)
	}
	return true, nil
}

func (c *likeCache) RemVideoLikeFromUser(ctx context.Context, key string, member interface{}) error {
	if err := c.client.ZRem(ctx, key, member).Err(); err != nil {
		return errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to remove redis set member: %v", err)
	}
	return nil
}

func (c *likeCache) AddVideoLikeToUser(ctx context.Context, key string, score float64, member interface{}) error {
	if err := c.client.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: member,
	}).Err(); err != nil {
		return errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to add member to redis set: %v", err)
	}
	return nil
}

func (c *likeCache) GetVideoIdFromUserLike(ctx context.Context, key string, cursor, limit int64) ([]int64, error) {
	var z []redis.Z
	var err error
	if cursor == 0 {
		z, err = c.client.ZRevRangeWithScores(ctx, key, 0, limit-1).Result()
	} else {
		z, err = c.client.ZRevRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
			Max:   fmt.Sprintf("(%d", cursor),
			Count: limit,
		}).Result()
	}

	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis: falied to get video member list: %v", err)
	}

	l := len(z)
	ids := make([]int64, l)
	for i, v := range z {
		idStr, ok := v.Member.(string)
		if !ok {
			return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to get id")
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to parse id")
		}
		ids[i] = id
	}
	return ids, nil
}

func (c *likeCache) GetVideoLikeCount(ctx context.Context, key string) (int64, error) {
	total, err := c.client.ZCard(ctx, key).Result()
	if err != nil {
		return 0, errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to get user video like count: %v", err)
	}
	return total, nil
}

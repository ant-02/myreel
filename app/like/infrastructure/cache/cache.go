package cache

import (
	"context"
	"fmt"
	"myreel/app/like/domain/repository"
	"myreel/pkg/errno"

	"github.com/redis/go-redis/v9"
)

type likeCache struct {
	client *redis.Client
}

func NewVideoCache(client *redis.Client) repository.LikeCache {
	return &likeCache{client: client}
}

func (c *likeCache) IsExist(ctx context.Context, key string, val interface{}) (bool, error) {
	exist, err := c.client.SIsMember(ctx, key, val).Result()
	if err != nil {
		return false, errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to check redis set member: %v", err)
	}
	return exist, nil
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
	var members []redis.Z
	var err error
	if cursor == 0 {
		members, err = c.client.ZRevRangeWithScores(ctx, key, 0, limit-1).Result()
		if err != nil {
			return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis: falied to get video member list: %v", err)
		}
	} else {
		members, err = c.client.ZRevRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
			Max:   fmt.Sprintf("(%d", cursor),
			Count: limit,
		}).Result()
	}

	l := len(members)
	ids := make([]int64, l)
	for i, member := range members {
		id, ok := member.Member.(int64)
		if !ok {
			return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to parse id: %v", err)
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

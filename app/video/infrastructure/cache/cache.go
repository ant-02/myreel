package cache

import (
	"context"
	"fmt"
	"myreel/app/video/domain/model"
	"myreel/app/video/domain/repository"
	"myreel/app/video/infrastructure/cache/pack"
	"myreel/pkg/errno"
	"time"

	"github.com/redis/go-redis/v9"
)

type videoCache struct {
	client *redis.Client
}

func NewVideoCache(client *redis.Client) repository.VideoCache {
	return &videoCache{client: client}
}

func (vc *videoCache) IsExist(ctx context.Context, key string) bool {
	return vc.client.Exists(ctx, key).Val() == 1
}

func (vc *videoCache) AddPopularVideoId(ctx context.Context, key string, score float64, member interface{}) error {
	if err := vc.client.ZAdd(ctx, key, redis.Z{
		Score:  score,
		Member: member,
	}).Err(); err != nil {
		return errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to add popular video")
	}
	return nil
}

func (vs *videoCache) GetPopularVideos(ctx context.Context, key string, cursor, limit int64) ([]int64, error) {
	var members []redis.Z
	var err error
	if cursor == 0 {
		members, err = vs.client.ZRevRangeWithScores(ctx, key, 0, limit-1).Result()
		if err != nil {
			return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis: falied to get popular videos: %v", err)
		}
	} else {
		members, err = vs.client.ZRevRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
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

func (vs *videoCache) CleanPopularVideos(ctx context.Context, key string, limit int64) error {
	if err := vs.client.ZRemRangeByRank(ctx, key, limit, -1).Err(); err != nil {
		return errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to clean popular videos: %v", err)
	}
	return nil
}

func (vs *videoCache) AddVideoWithTLL(ctx context.Context, key string, val interface{}, ttl time.Duration) error {
	err := vs.client.HSet(ctx, key, val).Err()
	if err != nil {
		return errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to add video: %v", err)
	}

	err = vs.client.Expire(ctx, key, ttl).Err()
	if err != nil {
		return errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to add ttl: %v", err)
	}

	return nil
}

func (vs *videoCache) GetVideo(ctx context.Context, key string) (*model.Video, error) {
	val, err := vs.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to get video by id: %v", err)
	}
	video, err := pack.MapToVideo(val)
	if err != nil {
		return nil, err
	}
	return video, nil
}

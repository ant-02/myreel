package cache

import (
	"context"
	"fmt"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
	"strconv"
	"strings"
)

func (c *userCache) SetUserBaned(ctx context.Context, uid int64) error {
	key := fmt.Sprintf("%s%d", constants.RedisUserBanedKey, uid)
	err := c.client.Set(ctx, key, "1", constants.NeverExpire).Err()
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, err.Error())
	}
	return nil
}

func (c *userCache) DeleteUserBaned(ctx context.Context, uid int64) error {
	key := fmt.Sprintf("%s%d", constants.RedisUserBanedKey, uid)
	err := c.client.Del(ctx, key).Err()
	if err != nil {
		return errno.NewErrNo(errno.InternalRedisErrorCode, err.Error())
	}
	return nil
}

func (c *userCache) GetBannedUserIDs(ctx context.Context, pattern string) ([]int64, error) {
	keys, err := c.client.Keys(ctx, constants.RedisUserBanedKey+"*").Result()
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, err.Error())
	}

	var userIDs []int64
	for _, key := range keys {
		// 从key中提取userID
		if uidStr := strings.TrimPrefix(key, constants.RedisUserBanedKey); uidStr != "" {
			if uid, err := strconv.ParseInt(uidStr, 10, 64); err == nil {
				userIDs = append(userIDs, uid)
			}
		}
	}
	return userIDs, nil
}

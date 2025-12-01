package cache

import (
	"context"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
)

func (c *userCache) SetUserLogin(ctx context.Context, key string, token string) error {
	err := c.client.Set(ctx, key, token, constants.RefreshTokenTTL).Err()
	if err != nil {
		return errno.Errorf(errno.InternalRedisErrorCode, "userCache.SetUserLogin failed, %v", err)
	}
	return nil
}

func (c *userCache) DeleteUserLogin(ctx context.Context, key string) error {
	err := c.client.Del(ctx, key).Err()
	if err != nil {
		return errno.Errorf(errno.InternalRedisErrorCode, "userCache.DeleteUserLogin failed, %v", err)
	}
	return nil
}

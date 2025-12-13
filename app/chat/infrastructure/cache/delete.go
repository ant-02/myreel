package cache

import (
	"context"
	"myreel/pkg/errno"
	"strconv"
)

func (cc *chatCache) RemUnreadMessage(ctx context.Context, key string, vals ...int64) error {
	l := len(vals)
	ids := make([]string, l)
	if l > 0 {
		for i, v := range vals {
			ids[i] = strconv.FormatInt(v, 10)
		}
	}
	if err := cc.client.ZRem(ctx, key, ids).Err(); err != nil {
		return errno.Errorf(errno.InternalRedisErrorCode, "redis: failed to remove unread message by member: %v", err)
	}
	return nil
}

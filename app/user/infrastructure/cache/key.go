package cache

import (
	"fmt"
	"myreel/pkg/constants"
)

func (c *userCache) UserBanedKey(uid int64) string {
	return fmt.Sprintf(constants.RedisUserBanedKey+"%d", uid)
}

func (c *userCache) UserLoginKey(uid int64) string {
	return fmt.Sprintf(constants.RedisUserLoginKey+"%d", uid)
}

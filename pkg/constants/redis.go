package constants

import "time"

const (
	RedisSlowQuery = 10 // ms redis默认的慢查询时间，适用于 logger
)

// Redis Key and Expire Time
const (
	PingTime                 = 2 * time.Second
	RedisCartExpireTime      = 5 * 60 * time.Second
	RedisCartStoreNum        = 30
	RedisSpuImageExpireTime  = 5 * 60 * time.Second
	RedisSkuImagesExpireTime = 5 * 60 * time.Second
	RedisLockStockExpireTime = 24 * 60 * 60 * time.Second
	RedisStockExpireTime     = 24 * 60 * 60 * time.Second
	RedisNXExpireTime        = 3 * time.Second
	RedisMaxLockRetryTime    = 400 * time.Millisecond
	RedisRetryStopTime       = 100 * time.Millisecond
	NeverExpire              = 0

	RedisUserLoginExpireTime = 2 * 60 * 60 * time.Second

	RedisVideoExpireTime = 10 * time.Minute
)

// Redis Count
const (
	RedisVideoPopCount = 200
)

// Redis Keys
const (
	RedisUserLoginKey = "login:user:"
	RedisUserBanedKey = "ban:user:"

	RedisVideoPopularKey = "pop:video"
	RedisVideoKey        = "video:"
)

// Redis DB Name
const (
	RedisDBOrder     = 0
	RedisDBCommodity = 1
	RedisDBCart      = 2
	RedisDBGateWay
	RedSyncDBId = 0
)

// Redis Connection Pool Configuration
const (
	RedisPoolSize           = 50              // 最大连接数
	RedisMinIdleConnections = 10              // 最小空闲连接数
	RedisDialTimeout        = 5 * time.Second // 连接超时时间
)

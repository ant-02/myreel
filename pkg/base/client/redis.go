package client

import (
	"context"
	"errors"
	"fmt"
	"myreel/config"
	"myreel/pkg/constants"
	"myreel/pkg/errno"
	"myreel/pkg/logger"

	"github.com/redis/go-redis/v9"
)

func InitRedis(db int) (*redis.Client, error) {
	if config.Redis == nil {
		return nil, errors.New("redis config is nil")
	}

	client := redis.NewClient(&redis.Options{
		Addr:         config.Redis.Addr,
		Password:     config.Redis.Password,
		DB:           db,
		PoolSize:     constants.RedisPoolSize,           // 连接池大小
		MinIdleConns: constants.RedisMinIdleConnections, // 最小空闲连接数
		DialTimeout:  constants.RedisDialTimeout,        // 连接超时
	})

	// 添加日志 Hook
	l := logger.GetRedisLogger()
	redis.SetLogger(l)
	client.AddHook(l)

	// 使用超时的 Ping
	ctx, cancel := context.WithTimeout(context.Background(), constants.PingTime)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, fmt.Sprintf("client.NewRedisClient: ping redis failed: %v", err))
	}

	return client, nil
}

func NewRedisClient(db int) (*redis.Client, error) {
	if config.Redis == nil {
		return nil, errors.New("redis config is nil")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       db,
	})
	l := logger.GetRedisLogger()
	redis.SetLogger(l)
	client.AddHook(l)
	_, err := client.Ping(context.TODO()).Result()
	if err != nil {
		return nil, errno.NewErrNo(errno.InternalRedisErrorCode, fmt.Sprintf("client.NewRedisClient: ping redis failed: %v", err))
	}
	return client, nil
}

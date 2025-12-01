package logger

import (
	"context"
	"fmt"
	"myreel/pkg/constants"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type RedisLogger struct{}

func (l *RedisLogger) Printf(ctx context.Context, template string, args ...interface{}) {
	control.info(fmt.Sprintf(template, args...), zap.String(constants.SourceKey, constants.RedisSource))
}

func (l *RedisLogger) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}

func (l *RedisLogger) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		start := time.Now().UnixMilli()

		if err := next(ctx, cmd); err != nil {
			return err
		}

		consume := time.Now().UnixMilli() - start
		if consume >= constants.RedisSlowQuery {
			Warn(fmt.Sprintf("slowly redis query. consume %d microsecond, query: %s", consume, cmd.String()),
				zap.String(constants.SourceKey, constants.RedisSource))
		}

		return nil
	}
}

func (l *RedisLogger) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		return next(ctx, cmds)
	}
}

func GetRedisLogger() *RedisLogger {
	return &RedisLogger{}
}

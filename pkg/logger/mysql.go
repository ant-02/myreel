package logger

import (
	"fmt"
	"myreel/pkg/constants"

	"go.uber.org/zap"
)

type MysqlLogger struct{}

func (l *MysqlLogger) Printf(template string, args ...interface{}) {
	control.info(fmt.Sprintf(template, args...), zap.String(constants.SourceKey, constants.MysqlSource))
}

func GetMysqlLogger() *MysqlLogger {
	return &MysqlLogger{}
}

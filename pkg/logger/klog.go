package logger

import (
	"context"
	"fmt"
	"io"
	"myreel/pkg/constants"

	"github.com/cloudwego/kitex/pkg/klog"
	"go.uber.org/zap"
)

type KlogLogger struct{}

func GetKlogLogger() *KlogLogger {
	return &KlogLogger{}
}

func (l *KlogLogger) Trace(v ...interface{}) {
	control.debug(fmt.Sprint(v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) Debug(v ...interface{}) {
	control.debug(fmt.Sprint(v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) Info(v ...interface{}) {
	control.info(fmt.Sprint(v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) Notice(v ...interface{}) {
	control.info(fmt.Sprint(v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) Warn(v ...interface{}) {
	control.warn(fmt.Sprint(v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) Error(v ...interface{}) {
	control.error(fmt.Sprint(v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) Fatal(v ...interface{}) {
	control.fatal(fmt.Sprint(v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) Tracef(format string, v ...interface{}) {
	control.debug(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) Debugf(format string, v ...interface{}) {
	control.debug(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) Infof(format string, v ...interface{}) {
	control.info(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) Noticef(format string, v ...interface{}) {
	control.info(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) Warnf(format string, v ...interface{}) {
	control.warn(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) Errorf(format string, v ...interface{}) {
	control.error(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) Fatalf(format string, v ...interface{}) {
	control.fatal(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	control.debug(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	control.debug(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	control.info(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	control.info(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	control.warn(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	control.error(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	control.fatal(fmt.Sprintf(format, v...), zap.String(constants.SourceKey, constants.KlogSource))
}

func (l *KlogLogger) SetLevel(klog.Level) {
}

func (l *KlogLogger) SetOutput(io.Writer) {
}

package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"time"
)

// SQLLogger 这是自定义的 GORM Logger
type SQLLogger struct {
	Logger *logrus.Logger
	Cfg    logger.Config
}

// 确保我们实现了 gorm.io/gorm/Logger.Interface
var _ logger.Interface = (*SQLLogger)(nil)

// LogMode 改变日志级别时被调用，可以按需实现
func (l *SQLLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.Cfg.LogLevel = level
	return &newLogger
}

// Info 这三个可以根据需要实现或留空
func (l *SQLLogger) Info(ctx context.Context, msg string, data ...interface{})  {}
func (l *SQLLogger) Warn(ctx context.Context, msg string, data ...interface{})  {}
func (l *SQLLogger) Error(ctx context.Context, msg string, data ...interface{}) {}

// Trace 关键方法：GORM 每执行一次 SQL，都会调用这里
func (l *SQLLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	// 1) 取出 trace_id（若你在 handler 里用 context.WithValue("trace_id", ...)）
	traceID, _ := ctx.Value("trace_id").(string)

	// 2) 处理错误、慢查询等逻辑
	switch {
	case err != nil:
		// 打印带有 error, trace_id
		l.Logger.WithFields(logrus.Fields{
			"trace_id": traceID,
			"elapsed":  elapsed,
			"rows":     rows,
		}).WithError(err).Errorf("SQL => %s", sql)
	// 如果你想要慢查询日志
	case elapsed > l.Cfg.SlowThreshold && l.Cfg.SlowThreshold != 0:
		l.Logger.WithFields(logrus.Fields{
			"trace_id": traceID,
			"elapsed":  elapsed,
			"rows":     rows,
		}).Warnf("SLOW SQL => %s", sql)
	default:
		// 正常 SQL 日志
		l.Logger.WithFields(logrus.Fields{
			"trace_id": traceID,
			"elapsed":  elapsed,
			"rows":     rows,
		}).Infof("SQL => %s", sql)
	}
}

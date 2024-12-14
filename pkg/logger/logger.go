package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LogWithContext 从 gin.Context 获取 trace_id 并返回带 trace_id 的日志 entry
func LogWithContext(ctx *gin.Context) *logrus.Entry {
	traceID, exists := ctx.Get("trace_id")
	if !exists {
		traceID = "unknown"
	}

	return Logger.WithField("trace_id", traceID)
}

// Info 记录 Info 级别日志，带 ctx 上下文
func Info(ctx *gin.Context, message string) {
	LogWithContext(ctx).Info(message)
}

// Infof 记录 Info 级别日志，带格式化和 ctx 上下文
func Infof(ctx *gin.Context, format string, args ...interface{}) {
	LogWithContext(ctx).Infof(format, args...)
}

// Error 记录 Error 级别日志，带 ctx 上下文
func Error(ctx *gin.Context, message string) {
	LogWithContext(ctx).Error(message)
}

// Errorf 记录 Error 级别日志，带格式化和 ctx 上下文
func Errorf(ctx *gin.Context, format string, args ...interface{}) {
	LogWithContext(ctx).Errorf(format, args...)
}

// Warn 记录 Warn 级别日志，带 ctx 上下文
func Warn(ctx *gin.Context, message string) {
	LogWithContext(ctx).Warn(message)
}

// Warnf 记录 Warn 级别日志，带格式化和 ctx 上下文
func Warnf(ctx *gin.Context, format string, args ...interface{}) {
	LogWithContext(ctx).Warnf(format, args...)
}

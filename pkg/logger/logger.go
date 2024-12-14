package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Logger 从 gin.Context 获取 trace_id 并返回带 trace_id 的日志 entry
func Logger(ctx *gin.Context) *logrus.Entry {
	traceID, exists := ctx.Get("trace_id")
	if !exists {
		traceID = "unknown"
	}

	return BaseLogger.WithField("trace_id", traceID)
}

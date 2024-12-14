package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"innoversepm-backend/pkg/logger"
	"strings"
	"time"
)

// LoggerMiddleware 创建一个Gin中间件，使用Logrus记录请求日志
func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := ctx.GetHeader("x-trace-id")
		if traceID == "" {
			traceID = strings.ReplaceAll(uuid.New().String(), "-", "")
		}
		ctx.Header("x-trace-id", traceID)
		ctx.Set("trace_id", traceID)

		logger.Logger(ctx).Infof(
			"access log: %s|%s|%s|%s.", ctx.ClientIP(), ctx.Request.Method, ctx.Request.URL.Path, traceID)
		startTime := time.Now()
		// 处理请求
		ctx.Next()
		// 记录请求日志
		logger.Logger(ctx).Infof("response log: %d|%d.", ctx.Writer.Status(), time.Since(startTime))
	}
}

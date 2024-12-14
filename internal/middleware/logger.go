package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"time"
)

// LoggerMiddleware 创建一个Gin中间件，使用Logrus记录请求日志
func LoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := ctx.GetHeader("x-trace-id")
		if traceID == "" {
			traceID = uuid.New().String()
		}
		ctx.Header("x-trace-id", traceID)

		startTime := time.Now()
		// 处理请求
		ctx.Next()
		// 记录请求日志
		duration := time.Since(startTime)
		entry := logger.WithFields(logrus.Fields{
			"status_code": ctx.Writer.Status(),
			"method":      ctx.Request.Method,
			"path":        ctx.Request.URL.Path,
			"query":       ctx.Request.URL.RawQuery,
			"ip":          ctx.ClientIP(),
			"user_agent":  ctx.Request.UserAgent(),
			"duration":    duration.Seconds(),
			"trace_id":    traceID,
		})
		if len(ctx.Errors) > 0 {
			entry.Error("请求处理错误")
		} else {
			entry.Info("请求处理完成")
		}
	}
}

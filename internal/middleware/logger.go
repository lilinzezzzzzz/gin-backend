package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

// LoggerMiddleware 创建一个Gin中间件，使用Logrus记录请求日志
func LoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 将Logger注入上下文
		c.Set("logger", logger)

		startTime := time.Now()

		// 处理请求
		c.Next()

		// 记录请求日志
		duration := time.Since(startTime)
		entry := logger.WithFields(logrus.Fields{
			"status_code": c.Writer.Status(),
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
			"query":       c.Request.URL.RawQuery,
			"ip":          c.ClientIP(),
			"user_agent":  c.Request.UserAgent(),
			"duration":    duration.Seconds(),
		})

		if len(c.Errors) > 0 {
			entry.Error("请求处理错误")
		} else {
			entry.Info("请求处理完成")
		}
	}
}

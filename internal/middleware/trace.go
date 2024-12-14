package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TraceMiddleware 生成并设置 trace_id
func TraceMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := ctx.GetHeader("x-trace-id")
		if traceID == "" {
			traceID = uuid.New().String()
		}

		ctx.Set("x-trace-id", traceID)
		ctx.Next()
	}
}

package infra

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang-backend/pkg/db"
	"gorm.io/gorm"
)

func NewDB(ctx *gin.Context) *gorm.DB {
	traceID := ctx.GetString("trace_id")
	// 先把 gin.Context 转为标准 context，或者直接 ctx.Request.Context()
	stdCtx := ctx.Request.Context()
	// 手动把 trace_id 塞进去
	stdCtx = context.WithValue(stdCtx, "trace_id", traceID)
	newDB := db.DB.WithContext(stdCtx)
	return newDB
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/routes"
	"innoversepm-backend/internal/setting"
	"log"
)

func main() {
	// 初始化配置
	setting.LoadConfig()

	// 初始化 Gin 引擎
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	// 注册路由
	routes.RegisterRoutes(r)

	// 从配置中读取端口号
	port := setting.Config.App.Port
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

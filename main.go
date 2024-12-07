package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/config"
	"innoversepm-backend/internal/routes"
)

func main() {
	// 初始化配置
	config.InitConfig()

	// 初始化 Gin 引擎
	r := gin.Default()

	// 注册路由
	routes.RegisterRoutes(r)

	// 从配置中读取端口号
	port := config.Config.GetInt("app.port")
	r.Run(fmt.Sprintf(":%d", port))
}

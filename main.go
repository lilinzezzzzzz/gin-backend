package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-backend/internal/middleware"
	"golang-backend/internal/routers"
	"golang-backend/internal/setting"
	"golang-backend/pkg/constants"
	"golang-backend/pkg/db"
	"golang-backend/pkg/logger"
	"golang-backend/pkg/redis"
	"golang-backend/pkg/snowflake"
	"log"

	"os"
	"strings"
)

func main() {
	env := strings.ToLower(os.Getenv("GO_ENV"))
	switch env {
	case constants.DevEnvVal, constants.LocalEnvVal:
		gin.SetMode(gin.DebugMode)
	case constants.TestEnvVal:
		gin.SetMode(gin.TestMode)
	case constants.ProdEnvVal:
		gin.SetMode(gin.ReleaseMode)
	default:
		log.Fatalf("Invalid environment specified: %s", env)
	}
	// 加载配置
	setting.LoadConfig(env)

	// 初始化日志
	logger.InitLogrus(env)

	// 初始化 MySQL
	dbCfg := setting.Config.MySQL
	db.InitMySQL(
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.DBName,
		dbCfg.MaxOpenConns,
		dbCfg.MaxIdleConns,
		dbCfg.ConnMaxLifetime)
	defer db.CloseMySQL()

	// 初始化 Redis
	redisCfg := setting.Config.Redis
	redis.InitRedis(redisCfg.Host, redisCfg.Port, redisCfg.Password, redisCfg.DB)
	defer redis.CloseRedis()

	// 初始化 雪花算法
	snowflake.InitSnowflake()

	// 初始化 Gin 引擎
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.AuthMiddleware())

	// 注册路由
	routers.RegisterRoutes(r)

	// 从配置中读取端口号
	port := setting.Config.App.Port
	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

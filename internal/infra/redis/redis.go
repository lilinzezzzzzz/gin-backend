package redis

import (
	"context"
	"fmt"
	"innoversepm-backend/pkg/logger"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"innoversepm-backend/internal/setting"
)

var (
	Client *redis.Client
)

func InitRedis(cfg *setting.AppConfig) {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
		// 可选：增添连接超时、读写超时等参数
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	// 创建带有超时的上下文（例如2秒），以防初始化过程卡住
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := pingRedis(ctx); err != nil {
		logger.BaseLogger.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Redis connected successfully")
}

func pingRedis(ctx context.Context) error {
	_, err := Client.Ping(ctx).Result()
	return err
}

// CloseRedis 关闭 Redis 连接
func CloseRedis() {
	if err := Client.Close(); err != nil {
		log.Printf("Failed to close Redis connection: %v", err)
	} else {
		log.Println("Redis connection closed successfully")
	}
}

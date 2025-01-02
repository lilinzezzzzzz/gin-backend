package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang-backend/pkg/logger"
	"log"
	"time"
)

var Client *redis.Client

func InitRedis(host, port, password string, db int) {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
		// 可选：增添连接超时、读写超时等参数
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	// 创建带有超时的上下文（例如2秒），以防初始化过程卡住
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := PingRedis(ctx); err != nil {
		logger.BaseLogger.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Redis connected successfully")
}

func PingRedis(ctx context.Context) error {
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

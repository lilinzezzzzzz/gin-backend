package infra

import (
	"github.com/go-redis/redis/v8"
	pkgRedis "golang-backend/pkg/redis"
)

func NewRedisClient() *redis.Client {
	return pkgRedis.Client
}

package redis

import (
	"errors"
	"github.com/gin-gonic/gin"
	"time"

	"github.com/go-redis/redis/v8"
)

// SetValue 在Redis中设置键值对，并支持过期时间
func SetValue(ctx *gin.Context, key string, value interface{}, expiration time.Duration) error {
	return Client.Set(ctx, key, value, expiration).Err()
}

// GetValue 从Redis中获取字符串类型的值
func GetValue(ctx *gin.Context, key string) (string, error) {
	val, err := Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// 键不存在时，Get返回redis.Nil
		return "", nil
	}
	return val, err
}

// DeleteKey 删除指定的键
func DeleteKey(ctx *gin.Context, key string) error {
	return Client.Del(ctx, key).Err()
}

// CheckKeyExists 检查键是否存在
func CheckKeyExists(ctx *gin.Context, key string) (bool, error) {
	count, err := Client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// IncrementCounter 将指定key的值（应为整数）+1，如果key不存在则从0开始
func IncrementCounter(ctx *gin.Context, key string) (int64, error) {
	return Client.Incr(ctx, key).Result()
}

// DecrementCounter 将指定key的值（应为整数）-1
func DecrementCounter(ctx *gin.Context, key string) (int64, error) {
	return Client.Decr(ctx, key).Result()
}

// ExpireKey 为指定key设置过期时间
func ExpireKey(ctx *gin.Context, key string, expiration time.Duration) (bool, error) {
	return Client.Expire(ctx, key, expiration).Result()
}

// SetHash 设置哈希类型的字段值
func SetHash(ctx *gin.Context, key, field string, value interface{}) error {
	return Client.HSet(ctx, key, field, value).Err()
}

// GetHashField 获取哈希类型某个字段的值
func GetHashField(ctx *gin.Context, key, field string) (string, error) {
	val, err := Client.HGet(ctx, key, field).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	return val, err
}

// GetAllHash 获取哈希所有字段值
func GetAllHash(ctx *gin.Context, key string) (map[string]string, error) {
	return Client.HGetAll(ctx, key).Result()
}

// PushToList 向列表左侧插入元素
func PushToList(ctx *gin.Context, key string, values ...interface{}) (int64, error) {
	return Client.LPush(ctx, key, values...).Result()
}

// PopFromList 从列表左侧弹出元素
func PopFromList(ctx *gin.Context, key string) (string, error) {
	val, err := Client.LPop(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	return val, err
}

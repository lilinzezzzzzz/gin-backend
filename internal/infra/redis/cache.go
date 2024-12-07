package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

// SetValue 在Redis中设置键值对，并支持过期时间
func SetValue(key string, value interface{}, expiration time.Duration) error {
	return Client.Set(Ctx, key, value, expiration).Err()
}

// GetValue 从Redis中获取字符串类型的值
func GetValue(key string) (string, error) {
	val, err := Client.Get(Ctx, key).Result()
	if err == redis.Nil {
		// 键不存在时，Get返回redis.Nil
		return "", nil
	}
	return val, err
}

// DeleteKey 删除指定的键
func DeleteKey(key string) error {
	return Client.Del(Ctx, key).Err()
}

// CheckKeyExists 检查键是否存在
func CheckKeyExists(key string) (bool, error) {
	count, err := Client.Exists(Ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// IncrementCounter 将指定key的值（应为整数）+1，如果key不存在则从0开始
func IncrementCounter(key string) (int64, error) {
	return Client.Incr(Ctx, key).Result()
}

// DecrementCounter 将指定key的值（应为整数）-1
func DecrementCounter(key string) (int64, error) {
	return Client.Decr(Ctx, key).Result()
}

// ExpireKey 为指定key设置过期时间
func ExpireKey(key string, expiration time.Duration) (bool, error) {
	return Client.Expire(Ctx, key, expiration).Result()
}

// SetHash 设置哈希类型的字段值
func SetHash(key, field string, value interface{}) error {
	return Client.HSet(Ctx, key, field, value).Err()
}

// GetHashField 获取哈希类型某个字段的值
func GetHashField(key, field string) (string, error) {
	val, err := Client.HGet(Ctx, key, field).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

// GetAllHash 获取哈希所有字段值
func GetAllHash(key string) (map[string]string, error) {
	return Client.HGetAll(Ctx, key).Result()
}

// PushToList 向列表左侧插入元素
func PushToList(key string, values ...interface{}) (int64, error) {
	return Client.LPush(Ctx, key, values...).Result()
}

// PopFromList 从列表左侧弹出元素
func PopFromList(key string) (string, error) {
	val, err := Client.LPop(Ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

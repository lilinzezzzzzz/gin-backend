package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/entity"
	"innoversepm-backend/pkg/logger"
	"time"

	"github.com/go-redis/redis/v8"
)

func SessionCacheKey(session string) string {
	return "session:" + session
}

func SessionLstCacheKey(userID int64) string {
	return fmt.Sprintf("session_list:%d", userID)
}

// SetSession 设置会话键值，并设置过期时间（默认10800秒=3小时）
func SetSession(ctx *gin.Context, session string, userID int, category string, ex time.Duration) error {
	key := SessionCacheKey(session)
	value := map[string]interface{}{
		"user_id":  userID,
		"category": category,
	}
	return SetValue(ctx, key, value, ex)
}

// GetSessionValue  获取会话中的用户ID和用户类型。
func GetSessionValue(ctx *gin.Context, session string) (*entity.UserSessionData, error) {
	sessCacheKey := SessionCacheKey(session)
	value, err := GetValue(ctx, sessCacheKey)
	if err != nil {
		return nil, err
	}

	if value == "" {
		return nil, errors.New("session not found")
	}

	var intermediate entity.UserSessionData
	if err := json.Unmarshal([]byte(value), &intermediate); err != nil {
		logger.Logger.Error(fmt.Sprintf("Failed to unmarshal JSON, value: %s", value), err)
		return nil, err
	}

	return &intermediate, nil
}

// SetSessionList 更新用户的session列表：
// 如果列表长度<3，直接rpush
// 如果列表>=3，lpop最旧session再rpush新session
func SetSessionList(ctx *gin.Context, userID int64, session string) error {
	cacheKey := SessionLstCacheKey(userID)

	sessionList, err := GetListAll(ctx, cacheKey)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("Failed to get list from %s", cacheKey), err)
		return err
	}

	lengthSessionList := len(sessionList)

	// 不需要获取新连接，go-redis是并发安全的，多次使用Client即可
	if lengthSessionList < 3 {
		// 列表长度<3，直接插入
		if err := Client.RPush(ctx, cacheKey, session).Err(); err != nil {
			logger.Logger.Error(fmt.Sprintf("Failed to rpush session into %s", cacheKey), err)
			return err
		}
	} else {
		// 长度>=3，弹出旧的session
		oldSession, err := Client.LPop(ctx, cacheKey).Result()
		if errors.Is(err, redis.Nil) {
			// 如果为空，不用特别处理，但逻辑上来说已经判断length>=3，不会出现nil
		} else if err != nil {
			logger.Logger.Error(fmt.Sprintf("Failed to lpop from %s", cacheKey), err)
			return err
		} else {
			// 这里可根据需求删除旧的session键值，如果需要的话
			// err := Client.Del(ctx, SessionCacheKey(oldSession)).Err()
			// if err != nil {
			//     loggerError(fmt.Sprintf("Failed to delete old session %s", oldSession), err)
			//     return err
			// }

			// 插入新的session
			if err := Client.RPush(ctx, cacheKey, session).Err(); err != nil {
				logger.Logger.Error(fmt.Sprintf("Failed to rpush new session into %s", cacheKey), err)
				return err
			}

			logger.Logger.Warning(
				fmt.Sprintf("Session list for user %d is full, popping and deleting oldest session: %s",
					userID, oldSession),
			)
		}
	}

	return nil
}

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

// GetListAll 获取列表所有元素
func GetListAll(ctx *gin.Context, key string) ([]string, error) {
	return Client.LRange(ctx, key, 0, -1).Result()
}

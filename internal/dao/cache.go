package dao

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang-backend/internal/entity"
	"golang-backend/internal/infra"
	"golang-backend/internal/utils/logger"
	"time"

	"github.com/go-redis/redis/v8"
)

type Cache struct {
	cli *redis.Client
}

func NewCache() *Cache {
	return &Cache{
		cli: infra.NewRedisClient(),
	}
}

func (c *Cache) SessionCacheKey(session string) string {
	return "session:" + session
}

func (c *Cache) SessionLstCacheKey(userID uint) string {
	return fmt.Sprintf("session_list:%d", userID)
}

// SetSession 设置会话键值，并设置过期时间（默认10800秒=3小时）
func (c *Cache) SetSession(ctx *gin.Context, session string, userJson string) error {
	key := c.SessionCacheKey(session)
	return c.SetValue(ctx, key, userJson, time.Second*3600)
}

// GetSessionValue  获取会话中的用户ID和用户类型。
func (c *Cache) GetSessionValue(ctx *gin.Context, session string) (*entity.UserSessionData, error) {
	sessCacheKey := c.SessionCacheKey(session)
	value, err := c.GetValue(ctx, sessCacheKey)
	if err != nil {
		return nil, err
	}

	if value == "" {
		return nil, errors.New("session not found")
	}

	var intermediate entity.UserSessionData
	if err := json.Unmarshal([]byte(value), &intermediate); err != nil {
		logger.Logger(ctx).Error(fmt.Sprintf("Failed to unmarshal JSON, value: %s", value), err)
		return nil, err
	}

	return &intermediate, nil
}

func (c *Cache) DeleteSessionLst(ctx *gin.Context, userID uint) error {
	return c.DeleteKey(ctx, c.SessionLstCacheKey(userID))
}

// SetSessionList 更新用户的session列表：
// 如果列表长度<3，直接rpush
// 如果列表>=3，lpop最旧session再rpush新session
func (c *Cache) SetSessionList(ctx *gin.Context, userID uint, session string) error {
	cacheKey := c.SessionLstCacheKey(userID)

	sessionList, err := c.GetListAll(ctx, cacheKey)
	if err != nil {
		logger.Logger(ctx).Error(fmt.Sprintf("Failed to get list from %s", cacheKey), err)
		return err
	}

	lengthSessionList := len(sessionList)

	// 不需要获取新连接，go-redis是并发安全的，多次使用Client即可
	if lengthSessionList < 3 {
		// 列表长度<3，直接插入
		if err := c.cli.RPush(ctx, cacheKey, session).Err(); err != nil {
			logger.Logger(ctx).Error(fmt.Sprintf("Failed to rpush session into %s", cacheKey), err)
			return err
		}
	} else {
		// 长度>=3，弹出旧的session
		oldSession, err := c.cli.LPop(ctx, cacheKey).Result()
		if errors.Is(err, redis.Nil) {
			// 如果为空，不用特别处理，但逻辑上来说已经判断length>=3，不会出现nil
		} else if err != nil {
			logger.Logger(ctx).Errorf("Failed to lpop from, err: %s", err)
			return err
		} else {
			// 插入新的session
			if err := c.cli.RPush(ctx, cacheKey, session).Err(); err != nil {
				logger.Logger(ctx).Errorf("Failed to rpush new session into, err: %s", err)
				return err
			}

			logger.Logger(ctx).Warnf("Session list for user %d is full, popping and deleting oldest session: %s", userID, oldSession)
		}
	}
	_, err = c.ExpireKey(ctx, cacheKey, time.Second*3600)
	if err != nil {
		logger.Logger(ctx).Warnf("set %s expire faile, err: %v", cacheKey, err)
	}
	return nil
}

// SetValue 在Redis中设置键值对，并支持过期时间
func (c *Cache) SetValue(ctx *gin.Context, key string, value interface{}, expiration time.Duration) error {
	return c.cli.Set(ctx, key, value, expiration).Err()
}

// GetValue 从Redis中获取字符串类型的值
func (c *Cache) GetValue(ctx *gin.Context, key string) (string, error) {
	val, err := c.cli.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// 键不存在时，Get返回redis.Nil
		return "", nil
	}
	return val, err
}

// DeleteKey 删除指定的键
func (c *Cache) DeleteKey(ctx *gin.Context, key string) error {
	return c.cli.Del(ctx, key).Err()
}

// CheckKeyExists 检查键是否存在
func (c *Cache) CheckKeyExists(ctx *gin.Context, key string) (bool, error) {
	count, err := c.cli.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// IncrementCounter 将指定key的值（应为整数）+1，如果key不存在则从0开始
func (c *Cache) IncrementCounter(ctx *gin.Context, key string) (int64, error) {
	return c.cli.Incr(ctx, key).Result()
}

// DecrementCounter 将指定key的值（应为整数）-1
func (c *Cache) DecrementCounter(ctx *gin.Context, key string) (int64, error) {
	return c.cli.Decr(ctx, key).Result()
}

// ExpireKey 为指定key设置过期时间
func (c *Cache) ExpireKey(ctx *gin.Context, key string, expiration time.Duration) (bool, error) {
	return c.cli.Expire(ctx, key, expiration).Result()
}

// SetHash 设置哈希类型的字段值
func (c *Cache) SetHash(ctx *gin.Context, key, field string, value interface{}) error {
	return c.cli.HSet(ctx, key, field, value).Err()
}

// GetHashField 获取哈希类型某个字段的值
func (c *Cache) GetHashField(ctx *gin.Context, key, field string) (string, error) {
	val, err := c.cli.HGet(ctx, key, field).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	return val, err
}

// GetAllHash 获取哈希所有字段值
func (c *Cache) GetAllHash(ctx *gin.Context, key string) (map[string]string, error) {
	return c.cli.HGetAll(ctx, key).Result()
}

// PushToList 向列表左侧插入元素
func (c *Cache) PushToList(ctx *gin.Context, key string, values ...interface{}) (int64, error) {
	return c.cli.LPush(ctx, key, values...).Result()
}

// PopFromList 从列表左侧弹出元素
func (c *Cache) PopFromList(ctx *gin.Context, key string) (string, error) {
	val, err := c.cli.LPop(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	return val, err
}

// GetListAll 获取列表所有元素
func (c *Cache) GetListAll(ctx *gin.Context, key string) ([]string, error) {
	return c.cli.LRange(ctx, key, 0, -1).Result()
}

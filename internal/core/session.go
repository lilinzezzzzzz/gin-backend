package core

import (
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/infra/redis"
	"innoversepm-backend/pkg/logger"
	"strconv"
)

// VerifySession 验证 session，并返回 user_id (int) 和验证结果 (bool)
// 若验证失败，返回的 user_id 为0，bool为false
func VerifySession(ctx *gin.Context, session string) (int, bool) {
	if session == "" {
		logger.Logger.Warn("Token verification failed: token not found")
		return 0, false
	}

	value, err := redis.GetSessionValue(ctx, session)
	if err != nil {
		logger.Logger.Warn("Token verification failed: error getting session value: %v\n", err)
		return 0, false
	}

	if value == nil {
		logger.Logger.Warn("Token verification failed: session not found")
		return 0, false
	}

	userIDStr, ok := value["user_id"]
	if !ok || userIDStr == "" {
		logger.Logger.Warn("Token verification failed: user_id not found")
		return 0, false
	}

	// 将user_id从字符串转换为int
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		logger.Logger.Warn("Token verification failed: user_id parse error: %v\n", err)
		return 0, false
	}

	userCategory, ok := value["category"]
	if !ok {
		logger.Logger.Warn("Token verification failed: category not found for user_id: %d\n", userID)
		return 0, false
	}

	if userCategory != "manager" {
		logger.Logger.Warn("Token verification failed: invalid user_category: %s, user_id: %d\n", userCategory, userID)
		return 0, false
	}

	sessionLstKey := redis.SessionLstCacheKey(userID)
	sessionLst, err := redis.GetListAll(ctx, sessionLstKey)
	if err != nil {
		logger.Logger.Warn("Token verification failed: error getting session list: %v\n", err)
		return 0, false
	}

	if sessionLst == nil {
		logger.Logger.Warn("Token verification failed: session list nil for user_id: %d\n", userID)
		return 0, false
	}

	found := false
	for _, s := range sessionLst {
		if s == session {
			found = true
			break
		}
	}
	if !found {
		logger.Logger.Warn("Token verification failed: session not found in session list, user_id: %d\n", userID)
		return 0, false
	}

	return userID, true
}

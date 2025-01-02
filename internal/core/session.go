package core

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang-backend/internal/dao"
	"golang-backend/internal/entity"
	"golang-backend/internal/utils/logger"
)

// GenerateSession 生成一个 uuid 字符串
func GenerateSession() string {
	u := uuid.New()                 // 返回一个 [16]byte 的 UUID
	return hex.EncodeToString(u[:]) // 直接对 16 字节进行 Hex 编码
}

// VerifySession checks if the given session is valid and returns the user data associated with it.
func VerifySession(ctx *gin.Context, session string) (*entity.UserSessionData, bool) {
	if session == "" {
		logger.Logger(ctx).Warn("Token verification failed: token not found")
		return nil, false
	}

	cache := dao.NewCache()
	userData, err := cache.GetSessionValue(ctx, session)
	if err != nil {
		logger.Logger(ctx).Warnf("Token verification failed: error getting session userData: %v\n", err)
		return nil, false
	}

	if userData == nil {
		logger.Logger(ctx).Warn("Token verification failed: session not found")
		return nil, false
	}

	sessionLstKey := cache.SessionLstCacheKey(userData.ID)
	sessionLst, err := cache.GetListAll(ctx, sessionLstKey)
	if err != nil {
		logger.Logger(ctx).Warnf("Token verification failed: error getting session list: %v\n", err)
		return nil, false
	}

	if sessionLst == nil {
		logger.Logger(ctx).Warnf("Token verification failed: session list nil for user_id: %d\n", userData.ID)
		return nil, false
	}

	found := false
	for _, s := range sessionLst {
		if s == session {
			found = true
			break
		}
	}
	if !found {
		logger.Logger(ctx).Warnf("Token verification failed: session not found in session list, user_id: %d\n", userData.ID)
		return nil, false
	}

	return userData, true
}

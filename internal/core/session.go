package core

import (
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/entity"
	"innoversepm-backend/internal/infra/redis"
	"innoversepm-backend/pkg/logger"
)

// VerifySession checks if the given session is valid and returns the user data associated with it.
func VerifySession(ctx *gin.Context, session string) (*entity.UserSessionData, bool) {
	if session == "" {
		logger.Logger(ctx).Warn("Token verification failed: token not found")
		return nil, false
	}

	userData, err := redis.GetSessionValue(ctx, session)
	if err != nil {
		logger.Logger(ctx).Warn("Token verification failed: error getting session userData: %v\n", err)
		return nil, false
	}

	if userData == nil {
		logger.Logger(ctx).Warn("Token verification failed: session not found")
		return nil, false
	}

	sessionLstKey := redis.SessionLstCacheKey(userData.ID)
	sessionLst, err := redis.GetListAll(ctx, sessionLstKey)
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
		logger.Logger(ctx).Warn("Token verification failed: session not found in session list, user_id: %d\n", userData.ID)
		return nil, false
	}

	return userData, true
}

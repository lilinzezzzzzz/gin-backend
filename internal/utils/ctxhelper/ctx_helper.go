package ctxhelper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang-backend/internal/entity"
	"golang-backend/internal/utils/logger"
)

// GetUserData 从 gin.Context 中获取用户数据
func GetUserData(ctx *gin.Context) (*entity.UserSessionData, error) {
	val, exists := ctx.Get("user_data")
	if !exists {
		logger.Logger(ctx).Warn("userData not found")
		return nil, errors.New("userData not found")
	}

	userData, ok := val.(*entity.UserSessionData)
	if !ok {
		logger.Logger(ctx).Warn("invalid userData type")
		return nil, errors.New("invalid userData type")
	}

	return userData, nil
}

// GetUserID 从 gin.Context 中获取用户 ID
func GetUserID(ctx *gin.Context) (int64, error) {
	userData, err := GetUserData(ctx)
	if err != nil {
		return 0, err
	}

	return userData.ID, nil
}

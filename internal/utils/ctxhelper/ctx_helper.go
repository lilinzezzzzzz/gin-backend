package ctxhelper

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang-backend/internal/entity"
)

// GetUserData 从 gin.Context 中获取用户数据
func GetUserData(ctx *gin.Context) (*entity.UserSessionData, error) {
	val, exists := ctx.Get("user_data")
	if !exists {
		return nil, errors.New("userData not found")
	}

	userData, ok := val.(*entity.UserSessionData)
	if !ok {
		return nil, errors.New("invalid userData type")
	}

	return userData, nil
}

// GetUserID 从 gin.Context 中获取用户 ID
func GetUserID(ctx *gin.Context) (uint, error) {
	userData, err := GetUserData(ctx)
	if err != nil {
		return 0, err
	}

	return userData.ID, nil
}

package entity

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UserSessionData struct {
	Account   string     `json:"account"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	ID        int64      `json:"id"`
	DeletedAt *time.Time `json:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
	Category  string     `json:"category"`
}

func GetUserData(ctx *gin.Context) *UserSessionData {
	val, exists := ctx.Get("userData")
	if !exists {
		// 如果不存在，说明上下文中未设置userData
		// 根据业务需求决定如何处理，例如返回错误
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "userData not found"})
		return nil
	}

	// 对val进行类型断言
	userData, ok := val.(*UserSessionData)
	if !ok {
		// 类型断言失败，表示实际存储的类型不是*entity.UserSessionData
		// 根据需求处理错误
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid userData type"})
		return nil
	}

	return userData
}

func GetUserID(ctx *gin.Context) int64 {
	return GetUserData(ctx).ID
}

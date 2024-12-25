package core

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"innoversepm-backend/pkg/logger"
)

// HashPassword 对密码进行哈希加密。
func HashPassword(ctx *gin.Context, password string) (string, error) {
	// 生成哈希密码，bcrypt 默认生成盐并加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Logger(ctx).Error("bcrypt.GenerateFromPassword err: ", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword 验证密码是否匹配。
func VerifyPassword(plainPassword, hashedPassword string) (bool, error) {
	// 使用 bcrypt.CompareHashAndPassword 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		logger.Logger(nil).Error("bcrypt.CompareHashAndPassword err: ", err)
		return false, err
	}
	return true, nil
}

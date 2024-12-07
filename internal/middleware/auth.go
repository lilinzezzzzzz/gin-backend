package middleware

import (
	"innoversepm-backend/internal/setting"
	"innoversepm-backend/pkg/xsignature"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 不需要认证的路径列表
var notAuthPaths = []string{
	"/auth/login",
	"/auth/register",
	"/docs",
	"/openapi.json",
}

// AuthMiddleware 鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path

		// 1. 跳过无需认证的路径
		for _, path := range notAuthPaths {
			if urlPath == path || strings.HasPrefix(urlPath, "/docs") {
				c.Next()
				return
			}
		}

		// 2. 验签逻辑
		if strings.HasPrefix(urlPath, "/openapi") {
			xSignature := c.GetHeader("X-Signature")
			xTimestamp := c.GetHeader("X-Timestamp")
			xNonce := c.GetHeader("X-Nonce")

			signature := xsignature.NewSignatureSrv(setting.Config)
			if !signature.VerifySignature(xSignature, xTimestamp, xNonce) {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid xsignature or timestamp"})
				c.Abort()
				return
			}
			c.Next()
			return
		}

		// 3. Session 校验
		session := c.GetHeader("Authorization")
		userID, ok := utils.VerifySession(session)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid or missing session"})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("userID", userID)

		c.Next()
	}
}

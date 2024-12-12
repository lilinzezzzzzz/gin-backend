package middleware

import (
	"innoversepm-backend/internal/core"
	"innoversepm-backend/internal/setting"
	"innoversepm-backend/pkg/constants"
	"innoversepm-backend/pkg/xsignature"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 不需要认证的路径列表
var notSessionAuthPaths = []string{
	"/auth/login",
	"/auth/register",
	"/docs",
	"/openapi.json",
	"/v1/auth/login_by_account",
	"/v1/auth/login_by_phone",
	"/v1/auth/verify_session",
}

// AuthMiddleware 鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		urlPath := ctx.Request.URL.Path

		// 1. 跳过无需认证的路径
		for _, path := range notSessionAuthPaths {
			if urlPath == path || strings.HasPrefix(urlPath, "/docs") {
				ctx.Next()
				return
			}
		}

		// 2. 验签逻辑
		if strings.HasPrefix(urlPath, "/openapi") {
			xSignature := ctx.GetHeader("X-Signature")
			xTimestamp := ctx.GetHeader("X-Timestamp")
			xNonce := ctx.GetHeader("X-Nonce")

			signature := xsignature.NewSignatureSrv(setting.Config)
			if !signature.VerifySignature(ctx, xSignature, xTimestamp, xNonce) {
				ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid xsignature or timestamp"})
				ctx.Abort()
				return
			}
			ctx.Next()
			return
		}

		// 3. Session 校验
		session := ctx.GetHeader("Authorization")
		userData, ok := core.VerifySession(ctx, session)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid or missing session"})
			ctx.Abort()
			return
		}

		if strings.HasPrefix(urlPath, "/v1") {
			ctx.Next()
		}

		if userData.Category != constants.UserManager {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid user category"})
			ctx.Abort()
		}

		// 将用户信息存储到上下文中
		ctx.Set("userData", userData)
		ctx.Next()
	}
}

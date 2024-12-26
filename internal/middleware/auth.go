package middleware

import (
	"fmt"
	"golang-backend/internal/core"
	"golang-backend/internal/setting"
	"golang-backend/internal/utils/resp"
	"golang-backend/pkg/xsignature"
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

			signature := xsignature.NewSignatureSrv(setting.Config.App.SecretKey, "h256", 18000)
			ok, err := signature.VerifySignature(xSignature, xTimestamp, xNonce)
			if err != nil {
				resp.InternalServerError(ctx, fmt.Sprintf("signature VerifySignature err: %v", err))
				return
			}
			if !ok {
				resp.UNAUTHORIZED(ctx, "invalid xsignature or timestamp")
				return
			}
			ctx.Next()
			return
		}

		// 3. Session 校验
		session := ctx.GetHeader("Authorization")
		userData, ok := core.VerifySession(ctx, session)
		if !ok {
			resp.UNAUTHORIZED(ctx, "invalid or missing session")
			return
		}

		// 将用户信息存储到上下文中
		ctx.Set("user_data", userData)
		ctx.Next()
		return
	}
}

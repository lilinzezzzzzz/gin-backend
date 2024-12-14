package routers

import (
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/controllers"
	"innoversepm-backend/internal/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	// 用户路由示例
	authGroup := r.Group("/auth")
	authGroup.Use(middleware.AuthMiddleware())
	authController := controllers.NewAuthController()
	{
		authGroup.GET("/user", authController.UserSessionData)
		authGroup.POST("/login", authController.ManagerLogin)
		authGroup.PUT("/logout", authController.LoginOut)
	}
}

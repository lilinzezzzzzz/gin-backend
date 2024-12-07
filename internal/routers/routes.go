package routers

import (
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/controllers"
	"innoversepm-backend/internal/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	// 用户路由示例
	userGroup := r.Group("/users")
	userGroup.Use(middleware.AuthMiddleware())
	userController := controllers.NewAuthController()
	{
		userGroup.GET("/", userController.GetUsers)
		userGroup.POST("/", userController.CreateUser)
	}
}

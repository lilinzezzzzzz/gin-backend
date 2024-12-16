package routers

import (
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	// 用户路由示例
	authGroup := r.Group("/auth")
	authController := controllers.NewAuthController()
	{
		authGroup.GET("/user", authController.UserSessionData)
		authGroup.POST("/login", authController.ManagerLogin)
		authGroup.PUT("/logout", authController.LoginOut)
	}
	userGroup := r.Group("/user_manager")
	userController := controllers.NewUserController()
	{
		userGroup.GET("/user/:user_id", userController.UserDetail)
	}
}

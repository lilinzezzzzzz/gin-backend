package routers

import (
	"github.com/gin-gonic/gin"
	"golang-backend/internal/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	helloGroup := r.Group("/hello")
	helloServer := controllers.NewHelloController()
	helloGroup.GET("", helloServer.Hello)

	// 鉴权
	authGroup := r.Group("/auth")
	authController := controllers.NewAuthController()
	{
		authGroup.GET("/me", authController.AuthMe)
		authGroup.POST("/login", authController.UserLogin)
		authGroup.PUT("/logout", authController.UserLoginOut)
	}
	// 用户
	userGroup := r.Group("")
	userController := controllers.NewUserController()
	{
		userGroup.GET("/user/:user_id", userController.UserDetail)
	}
}

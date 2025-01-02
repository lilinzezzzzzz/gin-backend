package routers

import (
	"github.com/gin-gonic/gin"
	"golang-backend/internal/servers"
)

func RegisterRoutes(r *gin.Engine) {
	helloGroup := r.Group("/hello")
	helloServer := servers.NewHelloServer()
	helloGroup.GET("", helloServer.Hello)

	// 鉴权
	authGroup := r.Group("/auth")
	authServer := servers.NewAuthServer()
	{
		authGroup.GET("/me", authServer.AuthMe)
		authGroup.POST("/login", authServer.UserLogin)
		authGroup.PUT("/logout", authServer.UserLoginOut)
	}
	// 用户
	userGroup := r.Group("")
	userServer := servers.NewUserServer()
	{
		userGroup.GET("/user/:user_id", userServer.UserDetail)
	}
}

package routes

import (
	"github.com/gin-gonic/gin"
	"innoversepm-backend/internal/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	// 用户路由示例
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.POST("/", controllers.CreateUser)
	}
}

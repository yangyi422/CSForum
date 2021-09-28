package router

import (
	"CSForum/controller/v1/user"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	// 声明用户路由组 并注册对应路由
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register", user.Register)
		UserRouter.GET("login", user.Login)
		UserRouter.GET("list", user.List)
	}
}

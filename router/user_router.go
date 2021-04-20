package router

import "github.com/gin-gonic/gin"

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("register")
	}
}

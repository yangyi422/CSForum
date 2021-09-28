package router

import (
	_ "CSForum/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter(c *gin.Engine) {

	// 声明一个swagger对应的路由
	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 对路由进行分组
	PrivateGroup := c.Group("")
	{
		// 初始化用户相关路由
		InitUserRouter(PrivateGroup)
	}
}

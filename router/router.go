package router

import (
	"github.com/gin-gonic/gin"
	_ "CSForum/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter(c *gin.Engine) {

	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	PrivateGroup := c.Group("")
	{
		InitUserRouter(PrivateGroup)
	}
}

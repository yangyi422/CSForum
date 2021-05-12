package initialization

import (
	"CSForum/router"

	"github.com/gin-gonic/gin"
)

func InitRouter(c *gin.Engine) {

	PrivateGroup := c.Group("")
	{
		router.InitUserRouter(PrivateGroup)
	}
}

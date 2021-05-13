package main

import (
	"CSForum/router"

	"github.com/gin-gonic/gin"
)

func main() {
	var Router = gin.Default()
	router.InitRouter(Router)
	Router.Run()
}

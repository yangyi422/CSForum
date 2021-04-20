package main

import (
	"CSForum/initialization"

	"github.com/gin-gonic/gin"
)

func main() {
	var Router = gin.Default()
	initialization.InitRouter(Router)
	Router.Run()
}

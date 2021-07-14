package main

import (
	"CSForum/router"
	"CSForum/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := util.InitTrans("zh"); err != nil {
		fmt.Printf("初始化翻译器, err:%v\n", err)
		return
	}

	var Router = gin.Default()
	router.InitRouter(Router)
	Router.Run()
}

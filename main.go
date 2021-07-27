package main

import (
	"CSForum/initialization/logger"
	"CSForum/router"
	"CSForum/util"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志
	logger.InitLogger()
	// 初始化校验工具
	if err := util.InitTrans("zh"); err != nil {
		logger.Errorf("初始化翻译器, err:%v\n", err)
		return
	}
	// 声明路由并初始化
	var Router = gin.Default()
	router.InitRouter(Router)
	// 启动服务
	Router.Run()
}

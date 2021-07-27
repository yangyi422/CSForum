package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  interface{} `json:"msg"`
}

// 定义响应码常量
const (
	SUCCESS         = 0     // 成功
	JSON_ERROR      = 10001 // Json解析错误
	CHECK_ERROR     = 10002 // 数据校验错误
	OPERATION_ERROR = 10003 // 操作错误
)

// 响应结果
func Send(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, data)
	return
}

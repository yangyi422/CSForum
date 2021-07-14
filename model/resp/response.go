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
func Result(code int, data interface{}, msg interface{}, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// 默认成功响应
func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

// 带消息的成功响应
func OkWithMessage(message interface{}, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

// 带数据的成功响应
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

// 带消息和数据的成功响应
func OkWithDetailed(data interface{}, message interface{}, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

// 默认失败响应
func Fail(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, "操作失败", c)
}

// 带消息的失败响应
func FailWithMessage(code int, message interface{}, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

// 带数据的失败响应
func FailWithData(code int, data interface{}, c *gin.Context) {
	Result(code, data, "操作失败", c)
}

// 带消息和数据的失败响应
func FailWithDetailed(code int, data interface{}, message interface{}, c *gin.Context) {
	Result(code, data, message, c)
}

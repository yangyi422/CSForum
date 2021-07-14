package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS         = 0
	JSON_ERROR      = 10001
	CHECK_ERROR     = 10002
	OPERATION_ERROR = 10003
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

// 默认成功返回
func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

// 带消息的成功返回
func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

// 带数据的成功返回
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

// 带消息和数据的成功返回
func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

// 默认失败返回
func Fail(code int, c *gin.Context) {
	Result(code, map[string]interface{}{}, "操作失败", c)
}

// 带消息的失败返回
func FailWithMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

// 带数据的失败返回
func FailWithData(code int, data interface{}, c *gin.Context) {
	Result(code, data, "操作失败", c)
}

// 带消息和数据的失败返回
func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}

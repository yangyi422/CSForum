package resp

// 注册结构体(请求接口-结构体)
type Register struct {
	Code    int         `json:"code"`    // 返回码
	Message interface{} `json:"message"` // 返回信息
}

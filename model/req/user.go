package req

// 注册结构体(请求接口-结构体)
type Register struct {
	Username  string `json:"username,omitempty"`   // 用户名
	Password  string `json:"password,omitempty"`   // 密码
	NickName  string `json:"nick_name,omitempty"`  // 昵称
	HeaderImg string `json:"header_img,omitempty"` // 头像
}

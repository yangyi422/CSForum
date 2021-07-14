package req

// 注册结构体(请求接口-结构体)
type Register struct {
	Username string `json:"username,omitempty" binding:"required" label:"用户名"` // 用户名
	Password string `json:"password,omitempty" binding:"required" label:"密码"`  // 密码
	NickName string `json:"nick_name,omitempty" binding:"required" label:"昵称"` // 昵称
}

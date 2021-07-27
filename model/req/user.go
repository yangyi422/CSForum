package req

// 注册结构体(请求接口-结构体)
type Register struct {
	Username string `json:"username" binding:"required" label:"用户名"` // 用户名
	Password string `json:"password" binding:"required" label:"密码"`  // 密码
	NickName string `json:"nick_name" binding:"required" label:"昵称"` // 昵称
	Birthday string `json:"birthday" binding:"" label:"生日"`          // 生日
}

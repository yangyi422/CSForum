package resp

// 注册结构体(请求接口-结构体)
type Register struct {
	Code    int         `json:"code"`    // 返回码
	Message interface{} `json:"message"` // 返回信息
}

// 登录结构体(请求接口-结构体)
type Login struct {
	Code    int         `json:"code"`    // 返回码
	Message interface{} `json:"message"` // 返回信息
}

// 获取用户列表(请求接口-结构体)
type List struct {
	Code    int         `json:"code"`    // 返回码
	Message interface{} `json:"message"` // 返回信息
	Data    []ListData  `json:"data"`    // 数据
}

// 获取用户列表(请求接口-结构体-数据)
type ListData struct {
	UserId      int64  `json:"user_id" gorm:"column:user_id;index:idx_user_id;comment:'用户ID'"`     // 用户ID
	Birthday    string `json:"birthday" gorm:"column:birthday;comment:'用户生日'"`                     // 用户生日
	Age         int    `json:"age" gorm:"column:age;comment:'用户年龄'"`                               // 用户年龄
	Username    string `json:"username" gorm:"column:username;index:idx_user_name;comment:'用户姓名'"` // 用户名
	NickName    string `json:"nick_name" gorm:"column:nick_name;comment:'用户昵称'"`                   // 用户昵称
	Phone       string `json:"phone" gorm:"column:phone;comment:'用户手机号'"`                          // 用户手机号
	Email       string `json:"email" gorm:"column:email;comment:'用户邮箱'"`                           // 用户邮箱
	Address     string `json:"address" gorm:"column:address;comment:'用户地址'"`                       // 用户地址
	Description string `json:"description" gorm:"column:description;comment:'用户描述'"`               // 用户描述
	Status      string `json:"status" gorm:"column:status;comment:'用户状态'"`                         // 用户状态(0:未激活)
	HeaderImg   string `json:"header_img" gorm:"column:header_img;comment:'用户头像'"`                 // 用户头像
}
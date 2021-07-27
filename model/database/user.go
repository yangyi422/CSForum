package database

import (
	"gorm.io/gorm"
)

// 用户结构体
type User struct {
	gorm.Model
	UserId      int64  `gorm:"column:user_id;index:idx_user_id;comment:'用户ID'"`    // 用户ID
	Birthday    string `gorm:"column:birthday;comment:'用户生日'"`                     // 用户生日
	Age         int    `gorm:"column:age;comment:'用户年龄'"`                          // 用户年龄
	Username    string `gorm:"column:username;index:idx_user_name;comment:'用户姓名'"` // 用户名
	NickName    string `gorm:"column:nick_name;comment:'用户昵称'"`                    // 用户昵称
	Phone       string `gorm:"column:phone;comment:'用户手机号'"`                       // 用户手机号
	Email       string `gorm:"column:email;comment:'用户邮箱'"`                        // 用户邮箱
	Password    string `gorm:"column:password;comment:'用户密码'"`                     // 用户密码
	Salt        string `gorm:"column:salt;comment:'密码-盐'"`                         // 密码-盐
	Address     string `gorm:"column:address;comment:'用户地址'"`                      // 用户地址
	Description string `gorm:"column:description;comment:'用户描述'"`                  // 用户描述
	Status      string `gorm:"column:status;comment:'用户状态'"`                       // 用户状态(0:未激活)
	HeaderImg   string `gorm:"column:header_img;comment:'用户头像'"`                   // 用户头像
}

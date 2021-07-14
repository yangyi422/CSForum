package user

import (
	"CSForum/initialization"
	"CSForum/model/database"
	"CSForum/model/req"
	"CSForum/model/resp"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// 声明一个注册结构体的字段
	var register = req.Register{}
	// 使用上面声明的字段来接收请求发送的JSON字段
	if err := c.ShouldBindJSON(&register); err != nil {
		resp.FailWithMessage(resp.JSON_ERROR, err.Error(), c)
		return
	}
	// 声明一个用户结构体,并初始化相关字段
	user := &database.User{
		UserId:      0,
		Birthday:    "",
		Age:         0,
		Username:    register.Username,
		NickName:    register.NickName,
		Phone:       "",
		Email:       "",
		Password:    "",
		Salt:        "",
		Address:     "",
		Description: "",
		Status:      "",
		HeaderImg:   register.HeaderImg,
	}
	// 使用gorm将上面声明的用户数据添加到数据库中
	if err := initialization.DB.Create(&user).Error; err != nil {
		resp.FailWithMessage(resp.OPERATION_ERROR, err.Error(), c)
		return
	}
	resp.OkWithMessage("用户注册成功", c)
}

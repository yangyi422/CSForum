package user

import (
	"CSForum/initialization"
	"CSForum/model/database"
	"CSForum/model/req"
	"CSForum/model/resp"
	"CSForum/util"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var register req.Register
	if err := c.ShouldBindJSON(&register); err != nil {
		resp.FailWithMessage(resp.JSON_ERROR, err.Error(), c)
		return
	}
	if err := util.Verify(register, util.RegisterVerify); err != nil {
		resp.FailWithMessage(resp.CHECK_ERROR, err.Error(), c)
		return
	}
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
	if err := initialization.DB.Create(&user).Error; err != nil {
		resp.FailWithMessage(resp.OPERATION_ERROR, err.Error(), c)
		return
	}
	resp.Ok(c)
}

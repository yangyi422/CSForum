package user

import (
	init_mysql "CSForum/initialization/mysql"
	"CSForum/model/database"
	"CSForum/model/req"
	"CSForum/model/resp"
	"CSForum/util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary 用户注册
// @Description
// @Tags user
// @Accept json
// @Produce json
// @Param req body req.Register true "请求参数"
// @Success 200 {object} resp.Register
// @Failure 400 {string} string "返回No Found"
// @Failure 500 {string} string "返回internal error"
// @Router /api/e/get_e_page [post]
func Register(c *gin.Context) {
	// 声明一个注册结构体的字段
	var register = req.Register{}
	// 使用上面声明的字段来接收请求发送的JSON字段
	if err := c.ShouldBindJSON(&register); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			resp.FailWithMessage(resp.JSON_ERROR, err.Error(), c)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		resp.FailWithMessage(resp.CHECK_ERROR, util.RemoveTopStruct(errs.Translate(util.Trans)), c)
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
		HeaderImg:   "",
	}
	// 使用gorm将上面声明的用户数据添加到数据库中
	if err := init_mysql.DB.Create(&user).Error; err != nil {
		resp.FailWithMessage(resp.OPERATION_ERROR, err.Error(), c)
		return
	}
	resp.OkWithMessage("用户注册成功", c)
}

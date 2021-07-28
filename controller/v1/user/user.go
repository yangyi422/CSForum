package user

import (
	"CSForum/initialization/logger"
	"CSForum/initialization/mysql"
	"CSForum/model/database"
	"CSForum/model/req"
	"CSForum/model/resp"
	"CSForum/util"
	"time"

	"github.com/gin-gonic/gin"
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
// @Router /user/register [post]
func Register(c *gin.Context) {
	// 声明注册结构体的字段
	var register_req = req.Register{}
	var response = resp.Register{}
	// 使用上面声明的字段来接收请求发送的JSON字段
	if err := c.ShouldBindJSON(&register_req); err != nil {
		logger.Errorf("json解析错误，错误为:%+v", util.CheckOutReqJsonErr(err, c))
		return
	}
	// 生成用户ID
	var userID = util.GetRequestID()
	// 获取年龄
	var age int
	if len(register_req.Birthday) != 0 {
		date, err := time.Parse("2006-01-02", register_req.Birthday)
		if err != nil {
			logger.Errorf("生日日期解析错误，错误为:%+v", err.Error())
			response.Code = resp.CHECK_ERROR
			response.Message = "生日日期解析错误"
			resp.Send(response, c)
			return
		}
		var year = date.Year()
		if year <= 0 {
			age = 0
		} else {
			var nowYear = time.Now().Year()
			age = nowYear - year
		}
	}
	// 密码加盐// 获取随机数作为盐
	var salt, password = util.AddSalt(util.GetStrRequestID(), register_req.Password)
	// 声明一个用户结构体,并初始化相关字段
	user := &database.User{
		UserId:      userID,
		Birthday:    register_req.Birthday,
		Age:         age,
		Username:    register_req.Username,
		NickName:    register_req.NickName,
		Phone:       "",
		Email:       "",
		Password:    password,
		Salt:        salt,
		Address:     "",
		Description: "",
		Status:      "",
		HeaderImg:   "",
	}
	// 使用gorm将上面声明的用户数据添加到数据库中
	if err := mysql.GetMysql().Create(&user).Error; err != nil {
		logger.Errorf("用户添加至数据库错误，错误为:%+v", err.Error())
		response.Code = resp.OPERATION_ERROR
		response.Message = "用户添加至数据库错误"
		resp.Send(response, c)
		return
	}
	response.Code = resp.SUCCESS
	response.Message = "用户注册成功"
	resp.Send(response, c)
}

// @Summary 用户登陆
// @Description
// @Tags user
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Success 200 {object} resp.Register
// @Failure 400 {string} string "返回No Found"
// @Failure 500 {string} string "返回internal error"
// @Router /user/login [get]
func Login(c *gin.Context) {
	// 声明注册结构体的字段
	var response = resp.Login{}
	// 获取Url中的参数
	var username = c.Query("username")
	var password = c.Query("password")
	// 校验用户名和密码
	if len(username) == 0 {
		response.Code = resp.CHECK_ERROR
		response.Message = "用户名为必填字段"
		resp.Send(response, c)
		return
	}
	if len(password) == 0 {
		response.Code = resp.CHECK_ERROR
		response.Message = "用户名为必填字段"
		resp.Send(response, c)
		return
	}
	// 获取用户名对应的用户信息
	var user = database.User{}
	if err := mysql.GetMysql().Where("username = ?", username).Find(&user).Error; err != nil {
		logger.Errorf("用户%s信息查询错误，错误为:%+v", username, err.Error())
		response.Code = resp.OPERATION_ERROR
		response.Message = "用户信息查询错误"
		resp.Send(response, c)
		return
	}
	// 校验密码
	_, oldPassword := util.AddSalt(user.Salt, password)
	if oldPassword != user.Password {
		response.Code = resp.CHECK_ERROR
		response.Message = "用户密码错误"
		resp.Send(response, c)
		return
	}
	response.Code = resp.SUCCESS
	response.Message = "用户登录成功"
	resp.Send(response, c)
}

// @Summary 获取用户列表
// @Description
// @Tags user
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Param status query string true "状态"
// @Success 200 {object} resp.List
// @Failure 400 {string} string "返回No Found"
// @Failure 500 {string} string "返回internal error"
// @Router /user/list [get]
func List(c *gin.Context) {
	// 声明注册结构体的字段
	var response = resp.List{}
	// 获取Url中的参数
	var username = c.Query("username")
	var status = c.Query("status")
	// 设置查询表名称
	var subQuery = mysql.GetMysql().Table("user")
	// 根据条件，添加查询条件
	if len(username) != 0 {
		subQuery = subQuery.Where("username = ?", username)
	}
	if len(status) != 0 {
		subQuery = subQuery.Where("status = ?", status)
	}
	// 获取用户列表
	var data = []resp.ListData{}
	if err := subQuery.Find(&data).Error; err != nil {
		logger.Errorf("查询用户信息错误，错误为:%+v", err.Error())
		response.Code = resp.OPERATION_ERROR
		response.Message = "查询用户信息错误"
		resp.Send(response, c)
		return
	}
	response.Code = resp.SUCCESS
	response.Message = "查询用户信息成功"
	response.Data = data
	resp.Send(response, c)
}

package controllers

import (
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"wenzhi.com/gin-ranking/models"
)

// 使用结构体来避免包内函数重名
type UserController struct {}

func (u UserController) Register(c *gin.Context) {
	// 接收用户名、密码、确认密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")
	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}
	if password != confirmPassword {
		ReturnError(c, 4001, "两次密码输入不一致")
		return
	}
	
	user, _:= models.GetUserInfoByUsername(username)
	if user.Id != 0 {
		ReturnError(c, 4001, "用户名已存在")
	}
	_, err := models.AddUser(username, EncryMd5((password)))
	if err != nil {
		ReturnError(c, 4001, "保存失败，请联系管理员")
		return
	}
	ReturnSuccess(c, 0, "注册成功", true, 1)
}


type UserApi struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (u UserController) Login(c *gin.Context) {
	// 接收用户名、密码
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}
	user, _ := models.GetUserInfoByUsername(username)
	if user.Id == 0 {
		ReturnError(c, 4004, "用户名或密码错误")
		return
	}
	if user.Password != EncryMd5(password) {
		ReturnError(c, 4004, "用户名或密码错误")
		return
	}
	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	session.Save()

	data := UserApi{Id: user.Id, Username: user.Username}
	ReturnSuccess(c, 0, "登录成功", data, 1)
}


// // 	"wenzhi.com/gin-ranking/pkg/logger"
// // "fmt"
// import (
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"wenzhi.com/gin-ranking/models"
// )

// func GetUserInfo(c *gin.Context) {
// 	ReturnSuccess(c, 0, "success", "user information", 1)
// }


// func (u UserController) GetUserInfo(c *gin.Context) {

// 	// http://127.0.0.1:9999/user/info/name/zhangsan
// 	// 请求参数 方式一：/info/:id
// 	idStr := c.Param("id")
// 	name := c.Param("name")
// 	// ReturnSuccess(c, 0, name, id, 1)

// 	id, _ := strconv.Atoi(idStr)
// 	user, _ := models.GetUserTest(id)
// 	ReturnSuccess(c, 0, name, user, 1)

// 	// ReturnSuccess(c, 0, "success", "user information", 1)
// }

// func (u UserController) AddUser(c *gin.Context) {
// 	username := c.DefaultPostForm("username", "wangwu")
// 	id, err := models.AddUser(username)
// 	if err != nil {
// 		ReturnError(c, 4002, "保存错误")
// 		return
// 	}
// 	ReturnSuccess(c, 0, "保存成功", id, 1)
// }

// func (u UserController) UpdateUser(c *gin.Context) {
// 	username := c.DefaultPostForm("username", "")
// 	idStr := c.DefaultPostForm("id", "")
// 	id, _ := strconv.Atoi(idStr)
// 	 models.UpdateUser(id, username)
// 	 ReturnSuccess(c, 0, "更新成功", true, 1)
// }

// func (u UserController) DeleteUser(c *gin.Context) {
// 	idStr := c.DefaultPostForm("id", "")
// 	id, _ := strconv.Atoi(idStr)
// 	err := models.DeleteUser(id)
// 	if err != nil {
// 		ReturnError(c, 4002, "删除错误")
// 		return
// 	}
// 	ReturnSuccess(c, 0, "删除成功", true, 1)
// }

// func (u UserController) GetUserListTest(c *gin.Context) {
// 	users, err := models.GetUserListTest()
// 	if err != nil {
// 		ReturnError(c, 4004, "没有相关数据")
// 		return
// 	}
// 	ReturnSuccess(c, 0, "获取用户列表成功", users, 1)
// }

// // func GetUserList(c *gin.Context) {
// // 	ReturnError(c, 4004, "failed 没有相关信息")
// // }

// func (u UserController) GetList(c *gin.Context) {
// 	// logger.Write("日志信息", "user")

// 	// defer func() {
// 	// 	if err := recover(); err!= nil {
// 	// 		fmt.Println("捕获异常recover panic:", err)
// 	// 	}
// 	// }()

// 	num1 := 1
// 	num2 := 0
// 	num3 := num1 / num2
// 	// ReturnError(c, 4004, "failed 没有相关信息list")
// 	ReturnError(c, 4004, num3)
// }
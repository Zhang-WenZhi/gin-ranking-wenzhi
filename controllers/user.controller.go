package controllers

// 	"wenzhi.com/gin-ranking/pkg/logger"
// "fmt"
import (
	"github.com/gin-gonic/gin"
)

// 使用结构体来避免包内函数重名
type UserController struct {}


// func GetUserInfo(c *gin.Context) {
// 	ReturnSuccess(c, 0, "success", "user information", 1)
// }


func (u UserController) GetUserInfo(c *gin.Context) {

	// http://127.0.0.1:9999/user/info/name/zhangsan
	// 请求参数 方式一：/info/:id
	id := c.Param("id")
	name := c.Param("name")
	ReturnSuccess(c, 0, name, id, 1)

	// ReturnSuccess(c, 0, "success", "user information", 1)
}

// func GetUserList(c *gin.Context) {
// 	ReturnError(c, 4004, "failed 没有相关信息")
// }

func (u UserController) GetList(c *gin.Context) {
	// logger.Write("日志信息", "user")

	// defer func() {
	// 	if err := recover(); err!= nil {
	// 		fmt.Println("捕获异常recover panic:", err)
	// 	}
	// }()

	num1 := 1
	num2 := 0
	num3 := num1 / num2
	// ReturnError(c, 4004, "failed 没有相关信息list")
	ReturnError(c, 4004, num3)
}
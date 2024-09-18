package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wenzhi.com/gin-ranking/controllers"
	"wenzhi.com/gin-ranking/pkg/logger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// 以路由 中间件 的形式，进行调用
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)

	user := r.Group("/user")
	{
		// user.GET("/info", controllers.GetUserInfo)
		// user.POST("/list", controllers.GetList)

		// 请求参数 方式一：/info/:id
		// user.GET("/info/:id/:name", controllers.UserController{}.GetUserInfo)
		user.GET("/info/:id", controllers.UserController{}.GetUserInfo)
		user.POST("/list", controllers.UserController{}.GetList)
		user.POST("/add", controllers.UserController{}.AddUser)
		user.POST("/update", controllers.UserController{}.UpdateUser)
		user.POST("/delete", controllers.UserController{}.DeleteUser)
		user.POST("/list/test", controllers.UserController{}.GetUserListTest)
		
		user.GET("/hello", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Hello world string")
		})

		// user.POST("/list", func(ctx *gin.Context) {
		// 	ctx.String(http.StatusOK, "user list")
		// })

		user.PUT("/add", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user add")
		})

		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "user delete")
		})
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}

	return r
}
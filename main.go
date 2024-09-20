package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	fmt.Println("Hello, World!")

// 	// get + json方式
// 	r := gin.Default()
// 	r.GET("/ping", func(ctx *gin.Context) {
// 		ctx.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	// get方式，带路径参数
// 	r.GET("/hello/:name", func(ctx *gin.Context) {
// 		name := ctx.Param("name")
// 		ctx.JSON(200, gin.H{
// 			"message": "Hello " + name,
// 		})
// 	})

// 	// post方式，json参数绑定
// 	r.POST("/login", func(ctx *gin.Context) {
// 		var user struct {
// 			Username string `json:"username" binding:"required"`
// 			Password string `json:"password" binding:"required"`
// 		}

// 		if err := ctx.ShouldBindJSON(&user); err != nil {
// 			ctx.JSON(400, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		if user.Username == "admin" && user.Password == "password" {
// 			ctx.JSON(200, gin.H{
// 				"message": "Login successful",
// 			})
// 		} else {
// 			ctx.JSON(401, gin.H{
// 				"error": "Invalid username or password",
// 			})
// 		}
// 	})

// 	// get+String方式
// 	// r.GET("/user/hello", func(ctx *gin.Context) {
// 	// 	ctx.String(http.StatusOK, "Hello world string")
// 	// })

// 	// r.POST("/user/list", func(ctx *gin.Context) {
// 	// 	ctx.String(http.StatusOK, "user list")
// 	// })

// 	// r.PUT("/user/add", func(ctx *gin.Context) {
// 	// 	ctx.String(http.StatusOK, "user add")
// 	// })

// 	// r.DELETE("/user/delete", func(ctx *gin.Context) {
// 	// 	ctx.String(http.StatusOK, "user delete")
// 	// })

// 	user := r.Group("/user")
// 	{
// 		user.GET("/hello", func(ctx *gin.Context) {
// 			ctx.String(http.StatusOK, "Hello world string")
// 		})

// 		user.POST("/list", func(ctx *gin.Context) {
// 			ctx.String(http.StatusOK, "user list")
// 		})

// 		user.PUT("/add", func(ctx *gin.Context) {
// 			ctx.String(http.StatusOK, "user add")
// 		})

// 		user.DELETE("/delete", func(ctx *gin.Context) {
// 			ctx.String(http.StatusOK, "user delete")
// 		})
// 	}

// 	//r.Run()
// 	// 指定端口号
// 	r.Run(":9999")

// }

import (
	"wenzhi.com/gin-ranking/router"
)

func main() {
	r := router.Router()
	r.Run(":9999")
	// Mac 环境下 打包：GOOS=linux GOARCH=amd64 go build
}

// import (
// 	"fmt"
// 	"wenzhi.com/gin-ranking/router"
// )

// func main() {
// 	r := router.Router()

// 	// go: 延迟捕获异常的方式
// 	// defer recover panic nil
// 	// 先defer的后执行，后defer的先执行

// 	defer func() {
// 		if err := recover(); err!= nil {
// 			fmt.Println("捕获异常recover panic:", err)
// 		}
// 	}

// 	defer fmt.Println(1)
// 	defer fmt.Println(2)
// 	defer fmt.Println(3)
// 	// 让程序崩溃
// 	panic("11")

// 	r.Run(":9999")
// }

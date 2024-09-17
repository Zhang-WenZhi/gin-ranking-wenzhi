package controllers

import (
	"github.com/gin-gonic/gin"
)

type OrderController struct{}

type Search struct {
	Name string `json:"name"` // 注意必须指定字段，否则的话首字母大写匹配不到的
	Cid  int    `json:"cid"`
}

// 注意没有*号
func (o OrderController) GetList(c *gin.Context) {
	// 请求参数 方式二：form
	// cid := c.PostForm("cid")
	// name := c.DefaultPostForm("name", "wangwu")
	// ReturnSuccess(c, 0, cid, name, 1)

	// 请求参数 方式三：json
	// params := make(map[string]interface{})
	// err := c.BindJSON(&params)
	// err := c.BindJSON(&search)
	// if err == nil {
	// 	ReturnSuccess(c, 0, "success", params, 1)
	// 	return
	// }
	// ReturnError(c, 4001, gin.H{"err": err})

	// 请求参数 方式三：j请求结构体的方式
	search := &Search{}
	err := c.BindJSON(&search)
	if err == nil {
		ReturnSuccess(c, 0, search.Name, search.Cid, 1)
		return
	}
	ReturnError(c, 4001, gin.H{"err": err})
	

	// ReturnError(c, 4004, "failed 没有相关信息order")
}





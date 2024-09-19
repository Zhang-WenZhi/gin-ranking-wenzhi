package controllers

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
	"github.com/gin-gonic/gin"
)


type JsonStruct struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Count int64       `json:"count"`
}

type JsonErrStruct struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// func ReturnSuccess1(c *gin.Context, data interface{}) {
// 	c.JSON(200, JsonStruct{
// 		Code: 0,
// 		Msg:  "success",
// 		Data: data,
// 		Count: 0,
// 	})
// }

// func ReturnError1(c *gin.Context, code int, msg string) {
// 	c.JSON(200, JsonStruct{
// 		Code: code,
// 		Msg:  msg,
// 		Data: nil,
// 		Count: 0,
// 	})
// }


func ReturnSuccess(c *gin.Context, code int, msg string, data interface{}, count int64) {
	json := &JsonStruct{
		Code: code,
		Msg:  msg,
		Data: data,
		Count: count,
	}
	c.JSON(200, json)
}

// func ReturnError(c *gin.Context, code int, msg string) {
// 	json := &JsonErrStruct{
// 		Code: code,
// 		Msg:  msg,
// 	}
// 	c.JSON(200, json)
// }

// func ReturnError(c *gin.Context, code int, msg interface{}) {
// 	strMsg, ok := msg.(string)
// 	if !ok {
// 			strMsg = "unknown error" // or handle the error appropriately
// 	}
// 	json := &JsonErrStruct{
// 			Code: code,
// 			Msg:  strMsg,
// 	}
// 	c.JSON(200, json)
// }


func ReturnError(c *gin.Context, code int, msg interface{}) {
	json := &JsonErrStruct{
			Code: code,
			Msg:  fmt.Sprintf("%v", msg),
	}
	c.JSON(200, json)
}

func EncryMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"wenzhi.com/gin-ranking/models"
)

type PlayerController struct {}

func (p PlayerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)
	rs, err := models.GetPlayers(aid, "id asc") // 按照id正序排列
	if err != nil {
		ReturnError(c, 4004, "获取玩家列表失败，没有相关信息")
		return
	}
	ReturnSuccess(c, 0, "success", rs, 1)

}

func (p PlayerController) GetRanking(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)
	rs, err := models.GetPlayers(aid, "score desc") // 按照分数倒序排列
	if err != nil {
		ReturnError(c, 4004, "获取排行榜失败，没有相关信息")
		return
	}
	ReturnSuccess(c, 0, "success", rs, 1)
	// redundant return statement (S1023)go-staticcheck
	// return
}
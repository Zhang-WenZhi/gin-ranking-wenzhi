package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"wenzhi.com/gin-ranking/cache"
	"wenzhi.com/gin-ranking/models"
)

type VoteController struct{}

func (v VoteController) AddVote(c *gin.Context) {
	userIdStr := c.DefaultPostForm("userId", "0")
	playerIdStr := c.DefaultPostForm("playerId", "0")
	userId, _ := strconv.Atoi(userIdStr)
	playerId, _ := strconv.Atoi(playerIdStr)
	if userId == 0 || playerId == 0 {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}
	user, _ := models.GetUserInfo(userId)
	if user.Id == 0 {
		ReturnError(c, 4001, "投票用户不存在")
		return
	}
	player, _ := models.GetPlayerInfo(playerId)
	if player.Id == 0 {
		ReturnError(c, 4001, "参赛选手不存在")
		return
	}
	vote, _ := models.GetVoteInfo(userId, playerId)
	if vote.Id != 0 {
		ReturnError(c, 4001, "您已经投过票了")
		return
	}

	rs, err := models.AddVote(userId, playerId)
	if err == nil {
		// 更新参赛选手的得票数，自增1
		models.UpdatePlayerScore(playerId)
		// 同时更新redis
		// var redisKey string // should merge variable declaration with assignment on next line (S1021)go-staticcheck ??
		// redisKey = "ranking:" + strconv.Itoa(player.Aid)
		var redisKey string = "ranking:" + strconv.Itoa(player.Aid)
		cache.Rdb.ZIncrBy(cache.Rctx, redisKey, 1, strconv.Itoa(playerId))

		ReturnSuccess(c, 0, "投票成功", rs, 1)
		return
	} else {
		ReturnError(c, 4004, "投票用户不存在")
		return
	}
	// ReturnError(c, 4004, "投票用户不存在")
	// redundant return statement (S1023)go-staticcheck
	// return
}
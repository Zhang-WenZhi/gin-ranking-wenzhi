package controllers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"wenzhi.com/gin-ranking/cache"
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

	// 测试redis缓存
	// err := cache.Rdb.Set(cache.Rctx, "name", "zhangsan", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	// var redisKey string
	// redisKey = "ranking:" + aidStr
	var redisKey string = "ranking:" + aidStr
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 1, -1).Result()
	if err == nil && len(rs) > 0 {
		var players []models.Player
		for _, value := range rs {
			id, _ := strconv.Atoi(value)
			rsInfo, _ := models.GetPlayerInfo(id)
			if rsInfo.Id > 0 {
				players = append(players, rsInfo)
			}
		}
		ReturnSuccess(c, 0, "success", players, 1)
		return
	}

	rsDb, errDb := models.GetPlayers(aid, "score desc") // 按照分数倒序排列
	if errDb == nil {
		for _, value := range rsDb {
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(value.Id, value.Score)).Err()
		}
		// 设置redis过期时间
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour)

		ReturnSuccess(c, 0, "success", rsDb, 1)
		return
	}
	ReturnError(c, 4004, "没有相关信息")
	// redundant return statement (S1023)go-staticcheck
	// return

	// ReturnSuccess(c, 0, "success", rsDb, 1)
	// redundant return statement (S1023)go-staticcheck
	// return
}
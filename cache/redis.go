package cache

import (
	"context"

	"github.com/redis/go-redis/v9"

	"wenzhi.com/gin-ranking/config"
)


var (
	Rdb *redis.Client
	Rctx context.Context
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: config.RedisAddress,
		Password: config.RedisPassword,
		DB: config.RedisDB,
	})
	Rctx = context.Background()
}

func Zscore(id int, score int) redis.Z {
	return redis.Z{Score: float64(score), Member: id}
}
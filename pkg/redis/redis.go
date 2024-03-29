package redis

import (
	"github.com/ClanceyLu/echo-api/conf"
	"github.com/go-redis/redis/v7"
)

// Client redis 实例
type Client = *redis.Client

// Nil 没有找到
const Nil = redis.Nil

// Connect 返回 redis 连接实例
func Connect() Client {
	var (
		redisConf = conf.Conf.Sub("redis")
		host      = redisConf.GetString("host")
	)
	return redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})
}

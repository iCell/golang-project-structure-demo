package infra

import (
	"context"
	"github.com/trustwallet/go-libs/cache/redis"
	"log"
)

func NewRedisConn() *redis.Redis {
	redis, err := redis.Init(context.Background(), "dsn")
	if err != nil {
		log.Fatalln(err)
	}
	return redis
}

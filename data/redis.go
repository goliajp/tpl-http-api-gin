package data

import (
	"github.com/goliajp/http-api-gin/env"
	"github.com/goliajp/rx/v2"
	"github.com/redis/go-redis/v9"
)

var Redis *rx.Redis

func GetRedis() *redis.Client {
	return Redis.Open()
}

func KvPrepare() {
	if env.KvRebuild {
		// do something
	}
	// do something
}

func init() {
	Redis = rx.NewRedis(
		&rx.RedisConfig{
			Host:     env.KvHost,
			Port:     env.KvPort,
			Password: env.KvPassword,
			Db:       env.KvDb,
		},
	)
}

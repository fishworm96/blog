package redis

import (
	"fmt"
	
	"blog/setting"

	"github.com/go-redis/redis"
)

var client *redis.Client

func Init(cfg *setting.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err = client.Ping().Result()
	return
}

func Close() {
	_ = client.Close()
}
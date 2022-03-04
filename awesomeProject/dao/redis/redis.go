package redis

import (
	"golangstudy/jike/awesomeProject/setttings"
	"strconv"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func Init(cfg *setttings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
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

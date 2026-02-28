package redis

import (
	"github.com/go-redis/redis/v8"
)

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}

package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}

func Ping(rdb *redis.Client) error {
	return rdb.Ping(context.Background()).Err()
}

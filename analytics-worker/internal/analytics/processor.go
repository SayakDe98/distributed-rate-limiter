package analytics

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Process(rdb *redis.Client, userID string, allowed bool) error {
	totalKey := fmt.Sprintf("analytics:%s:total", userID)
	blockedKey := fmt.Sprintf("analytics:%s:blocked", userID)

	rdb.Incr(ctx, totalKey)

	if !allowed {
		rdb.Incr(ctx, blockedKey)
	}

	return nil
}

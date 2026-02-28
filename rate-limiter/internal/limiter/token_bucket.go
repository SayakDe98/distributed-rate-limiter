package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	maxTokens     = 10
	refillRate    = 10
	refillSeconds = 1
)

func Allow(rdb *redis.Client, userID string) bool {
	ctx := context.Background()

	tokenKey := "tokens:" + userID
	tsKey := "ts:" + userID

	now := time.Now().Unix()

	lastRefill, _ := rdb.Get(ctx, tsKey).Int64()
	tokens, _ := rdb.Get(ctx, tokenKey).Int()

	if lastRefill == 0 {
		rdb.Set(ctx, tsKey, now, 0)
		rdb.Set(ctx, tokenKey, maxTokens-1, 0)
		return true
	}

	elapsed := now - lastRefill
	if elapsed >= refillSeconds {
		tokens = min(maxTokens, tokens+int(elapsed)*refillRate)
		rdb.Set(ctx, tsKey, now, 0)
	}

	if tokens <= 0 {
		return false
	}

	rdb.Set(ctx, tokenKey, tokens-1, 0)
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package main

import (
	"rate-limiter/internal/api"
	"rate-limiter/internal/kafka"
	"rate-limiter/internal/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	redisClient := redis.NewClient()
	producer := kafka.NewProducer()

	handler := api.NewHandler(redisClient, producer)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/check", handler.CheckRateLimit)

	r.Run(":8080")
}

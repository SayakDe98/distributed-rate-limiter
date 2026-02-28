package api

import (
	"net/http"
	"rate-limiter/internal/limiter"
	"rate-limiter/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Handler struct {
	redis    *redis.Client
	producer KafkaProducer
}

type KafkaProducer interface {
	Publish(event model.RateEvent) error
}

func NewHandler(r *redis.Client, p KafkaProducer) *Handler {
	return &Handler{redis: r, producer: p}
}

func (h *Handler) CheckRateLimit(c *gin.Context) {
	var req struct {
		UserID string `json:"user_id"`
	}

	c.BindJSON(&req)

	allowed := limiter.Allow(h.redis, req.UserID)

	event := model.RateEvent{
		UserID:  req.UserID,
		Allowed: allowed,
	}

	h.producer.Publish(event)

	c.JSON(http.StatusOK, gin.H{
		"allowed": allowed,
	})
}

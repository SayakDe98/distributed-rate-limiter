package consumer

import (
	"analytics-worker/internal/analytics"
	"analytics-worker/internal/redis"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type RateEvent struct {
	UserID  string `json:"user_id"`
	Allowed bool   `json:"allowed"`
}

func Start() {
	rdb := redis.NewClient()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "rate-events",
		GroupID: "analytics-group",
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		var event RateEvent
		json.Unmarshal(msg.Value, &event)

		analytics.Process(rdb, event.UserID, event.Allowed)
	}
}

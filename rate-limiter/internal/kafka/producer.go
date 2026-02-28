package kafka

import (
	"context"
	"encoding/json"
	"rate-limiter/internal/model"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer() *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP("kafka:9092"),
			Topic:    "rate-events",
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (p *Producer) Publish(event model.RateEvent) error {
	data, _ := json.Marshal(event)
	return p.writer.WriteMessages(context.Background(),
		kafka.Message{Value: data})
}

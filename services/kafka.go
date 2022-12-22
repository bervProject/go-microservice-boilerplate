package services

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducerInterface interface {
	Publish(ctx context.Context, message string) error
}

type KafkaProducer struct{}

type KafkaRequest struct {
	Message string `json:"message"`
}

func NewProducer() KafkaProducerInterface {
	return &KafkaProducer{}
}

func (producer *KafkaProducer) Publish(ctx context.Context, message string) (err error) {
	publisher := &kafka.Writer{
		Addr:                   kafka.TCP("localhost:9092"),
		Topic:                  "test-a",
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	err = publisher.WriteMessages(ctx,
		kafka.Message{
			Key:     []byte("test"),
			Value:   []byte(message),
			Headers: []kafka.Header{{Key: "hk", Value: []byte("hv")}},
		},
	)
	if err != nil {
		log.Println("failed to write messages:", err)
		return
	}

	if err := publisher.Close(); err != nil {
		log.Println("failed to close writer:", err)
		return err
	}
	return
}

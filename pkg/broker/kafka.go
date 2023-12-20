package broker

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	writer *kafka.Writer
}

func NewPublisher(addr []string, autocreateTopics bool) *KafkaPublisher {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(addr...),
		AllowAutoTopicCreation: autocreateTopics,
	}

	return &KafkaPublisher{writer: w}
}

func (kp KafkaPublisher) Publish(msg []kafka.Message) error {
	if err := kp.writer.WriteMessages(context.Background(), msg...); err != nil {
		return fmt.Errorf("error writing messages to kafka %w", err)
	}

	if err := kp.writer.Close(); err != nil {
		return fmt.Errorf("error closing kafka writer %w", err)
	}

	return nil
}

package broker

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	autocreateTopics bool
	writer           *kafka.Writer
}

func NewPublisher(addr []string, autocreateTopics bool) *KafkaPublisher {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(addr...),
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: autocreateTopics,
	}

	return &KafkaPublisher{writer: w}
}

func (kp KafkaPublisher) Publish(msg []kafka.Message) error {
	if err := kp.writer.WriteMessages(context.Background(), msg...); err != nil {
		return err
	}

	if err := kp.writer.Close(); err != nil {
		return err
	}

	return nil
}

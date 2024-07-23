package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/yuxi311/webService/pkg/logger"
)

func ProduceMessage(msg []byte) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	w := KafkaWriter()
	
	err := w.WriteMessages(ctx, kafka.Message{Value: msg})
	if err != nil {
		logger.Errorf("failed to write messages, error: %v", err)
	}

	return nil
}

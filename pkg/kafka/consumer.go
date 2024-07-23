package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/yuxi311/webService/internal/config"
	"github.com/yuxi311/webService/pkg/logger"
)

func ConsumeMessage() error {
	cfg := config.Kafka()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{cfg.Brokers},
		Topic:    cfg.LoginLogTopic,
		GroupID:  "group1",
		MaxBytes: 10e6, // 10MB
	})
	r.SetOffset(kafka.LastOffset)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			break
		}
		logger.Infof("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	defer func() {
		if err := r.Close(); err != nil {
			logger.Errorf("failed to close reader: %v", err)
		}
	}()

	return nil
}

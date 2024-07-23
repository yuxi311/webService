package kafka

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/yuxi311/webService/internal/config"
	"github.com/yuxi311/webService/pkg/logger"
)

var internal_kafka *kafka.Conn

func Init() error {
	cfg := config.Kafka()
	ctx := context.Background()
	network := "tcp"
	partition := 0

	conn, err := kafka.DialLeader(ctx, network, cfg.Brokers, cfg.LoginLogTopic, partition)
	if err != nil {
		logger.Errorf("failed to dial leader, error: %v", err)
		return err
	}
	internal_kafka = conn
	return nil
}

func KafkaConn() *kafka.Conn {
	return internal_kafka
}

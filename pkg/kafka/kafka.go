package kafka

import (
	"github.com/segmentio/kafka-go"
	"github.com/yuxi311/webService/internal/config"
)

var internal_kafka *kafka.Writer

func Init() error {
	cfg := config.Kafka()

	w := &kafka.Writer{
		Addr:     kafka.TCP(cfg.Brokers),
		Topic:    cfg.LoginLogTopic,
		Balancer: &kafka.LeastBytes{},
	}

	internal_kafka = w
	return nil
}

func KafkaWriter() *kafka.Writer {
	return internal_kafka
}

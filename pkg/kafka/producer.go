package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/yuxi311/webService/pkg/logger"
)

func ProduceMessage(msg []byte) error {
	conn := KafkaConn()

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteMessages(
		kafka.Message{Value: msg},
	)
	if err != nil {
		logger.Errorf("failed to write messages, error: %v", err)
		return err
	}

	return nil
}

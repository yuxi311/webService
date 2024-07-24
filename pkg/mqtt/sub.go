package mqtt

import (
	"fmt"

	"github.com/yuxi311/webService/internal/config"
)

func Sub() error {
	cfg := config.MQTT()
	client := MQTTClient()

	token := client.Subscribe(cfg.Topic, 0, messagePubHandler)
	token.Wait()

	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("sub token to topic= %v error, err = %v", cfg.Topic, token.Error())
	}

	return nil
}

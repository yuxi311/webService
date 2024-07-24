package mqtt

import (
	"fmt"

	"github.com/yuxi311/webService/internal/config"
)

func Pub(msg []byte) error {
	cfg := config.MQTT()
	client := MQTTClient()

	token := client.Publish(cfg.Topic, 0, false, msg)
	token.Wait()
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("pub token error, err = %v", token.Error())
	}

	return nil
}

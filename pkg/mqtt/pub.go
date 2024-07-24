package mqtt

import (
	"fmt"
)

func Pub(topic string, qos byte, msg []byte) error {
	client := MQTTClient()

	token := client.Publish(topic, qos, false, msg)
	token.Wait()
	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("pub token error, err = %v", token.Error())
	}

	return nil
}

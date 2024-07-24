package mqtt

import (
	"fmt"
)

func Sub(topic string, qos byte) error {
	client := MQTTClient()

	token := client.Subscribe(topic, qos, messagePubHandler)
	token.Wait()

	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("sub token to topic= %v error, err = %v", topic, token.Error())
	}

	return nil
}

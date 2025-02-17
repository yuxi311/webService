package run

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/yuxi311/webService/dal"
	"github.com/yuxi311/webService/internal/config"
	"github.com/yuxi311/webService/internal/server"
	"github.com/yuxi311/webService/pkg/kafka"
	"github.com/yuxi311/webService/pkg/logger"
	"github.com/yuxi311/webService/pkg/mqtt"
)

func Run(configFile string) error {
	if err := config.Load(configFile); err != nil {
		return errors.Wrap(err, "config.load")
	}

	//init logger
	if err := logger.Init(); err != nil {
		return errors.Wrap(err, "logger.init")
	}
	
	//init db
	if err := dal.Init(); err != nil {
		return err
	}

	//init redis db
	if err := dal.InitRedis(); err != nil {
		logger.Errorf("init redis failed, error: %v", err.Error())
	}

	//init kafka
	if err := kafka.Init(); err != nil {
		logger.Errorf("init kafka failed, error: %v", err.Error())
	} else {
		go kafka.ConsumeMessage()
	}

	//init mqtt
	if err := mqtt.Init(); err != nil {
		logger.Errorf("init mqtt failed, error: %v", err.Error())
	} else {
		mqttCfg := config.MQTT()
		go mqtt.Sub(mqttCfg.Topic, 0)
	}

	// listen port
	port := config.Server().Port
	logger.Info("Start listen port")
	fmt.Printf("Start listen port: %v\n", port)
	server.Serve(port)
	return nil
}

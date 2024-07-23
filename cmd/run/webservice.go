package run

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/yuxi311/webService/dal"
	"github.com/yuxi311/webService/internal/config"
	"github.com/yuxi311/webService/internal/server"
	"github.com/yuxi311/webService/pkg/kafka"
	"github.com/yuxi311/webService/pkg/logger"
	"github.com/yuxi311/webService/pkg/utils"
)

func Run(configFile string) error {
	if err := config.Load(configFile); err != nil {
		return errors.Wrap(err, "config.load")
	}

	//init db
	if err := dal.Init(); err != nil {
		return err
	}

	//init redis db
	if err := dal.InitRedis(); err != nil {
		return err
	}

	//init logger
	loggerOptions := logger.Options{
		Mode:       config.Log().Mode,
		Level:      config.Log().Level,
		Path:       utils.ToFilePath(config.Log().File),
		Format:     config.Log().Format,
		MaxSize:    config.Log().MaxSize,
		MaxBackups: config.Log().MaxBackups,
	}
	if err := logger.Init(loggerOptions); err != nil {
		return errors.Wrap(err, "logger.init")
	}

	//init kafka
	if err := kafka.Init(); err != nil {
		return err
	}

	go kafka.ConsumeMessage()

	// listen port
	port := config.Server().Port
	logger.Info("Start listen port")
	fmt.Printf("Start listen port: %v\n", port)
	server.Serve(port)
	return nil
}

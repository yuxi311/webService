package run

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/yuxi311/webService/dal"
	"github.com/yuxi311/webService/internal/config"
	"github.com/yuxi311/webService/internal/server"
	"github.com/yuxi311/webService/pkg/logger"
	"github.com/yuxi311/webService/pkg/utils"
)

func Run(configFile string) error {
	if err := config.Load(configFile); err != nil {
		return errors.Wrap(err, "config.load")
	}

	//数据库的连接
	if err := dal.Init(); err != nil {
		return err
	}

	//logger init
	loggerOptions := logger.Options{
		Mode:       config.Log().Mode,
		Level:      config.Log().Level,
		Path:       utils.ToFilePath(config.Log().File),
		Format:     config.Log().Format,
		MaxSize:    config.Log().MaxSize,
		MaxBackups: config.Log().MaxBackups,
	}
	if err := logger.Initialize(loggerOptions); err != nil {
		return errors.Wrap(err, "logger.init")
	}

	// listen port
	port := config.Server().Port
	logger.Info("Start listen port")
	fmt.Printf("Start listen port: %v\n", port)
	server.Serve(port)
	return nil
}

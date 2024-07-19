package run

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/yuxi311/webService/dal"
	"github.com/yuxi311/webService/internal/config"
	"github.com/yuxi311/webService/internal/server"
)

func Run(configFile string) error {
	if err := config.Load(configFile); err != nil {
		return errors.Wrap(err, "config.load")
	}

	//数据库的连接
	if err := dal.Init(); err != nil {
		return err
	}

	// listen port
	port := config.Server().Port
	fmt.Printf("Start listen port: %v\n", port)
	server.Serve(port)
	return nil
}

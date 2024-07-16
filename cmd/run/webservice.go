package run

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/yuxi311/webService/dal"
	"github.com/yuxi311/webService/internal/config"
	"github.com/yuxi311/webService/model/entity"
	"github.com/yuxi311/webService/server"
)

func Run(configFile string) error {
	if err := config.Load(configFile); err != nil {
		return errors.Wrap(err, "config.load")
	}

	//数据库的连接
	if err := dal.Init(); err != nil {
		return err
	}

	//数据库操作
	mysql := dal.DB()

	// 自动迁移模式
	mysql.AutoMigrate(&entity.User{})

	//创建记录
	user := entity.User{
		Name:     "web1",
		Username: "user1",
		Password: "0000",
		Role:     1,
	}
	mysql.Create(&user)

	//查询数据库
	var readUsers []entity.User
	mysql.Limit(3).Find(&readUsers)

	fmt.Printf("limit 1: %v\n", readUsers[0])
	fmt.Printf("limit 2: %v\n", readUsers[1])
	fmt.Printf("limit 3: %v\n", readUsers[2])

	//更新表中字段值
	mysql.Model(&readUsers[1]).Update("name", "web2")
	mysql.Model(&entity.User{}).Where("id = ?", 4).Update("name", "web3")

	//删除记录
	mysql.Where("name = ?", "web1").Delete(&entity.User{})

	port := config.Server().Port

	fmt.Printf("Start listen port: %v\n", port)
	server.Serve(port)
	return nil
}

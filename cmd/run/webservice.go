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

	//获取数据库连接对象
	// mysql := dal.DB()

	// // 自动迁移模式
	// dal.AutoMigrate()

	// //创建记录
	// user.CreateNewUser("web4", "user4", "0000", 1)

	// //查询数据库
	// queryUsers := user.QueryUsers()
	// fmt.Println("query users: ", queryUsers)

	// queryUser := user.QueryUser(4)
	// fmt.Println("query user: ", queryUser)

	// //更新表中字段值
	// user.UpdateUser(4, "", "1111", 1)

	// //删除记录
	// user.DeleteUser(3)
	// mysql.Where("name = ?", "web1").Delete(&entity.User{})

	// listen port
	port := config.Server().Port
	fmt.Printf("Start listen port: %v\n", port)
	server.Serve(port)
	return nil
}

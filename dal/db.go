package dal

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/yuxi311/webService/internal/config"
)

var internal_db *gorm.DB

func Init() error {
	cfg := config.DB()
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("open db error, err = %v", err)
	}

	if err := AutoMigrate(db); err != nil {
		return err
	}

	fmt.Println("connected database successful: ", db)
	internal_db = db
	return nil
}

func DB() *gorm.DB {
	return internal_db
}

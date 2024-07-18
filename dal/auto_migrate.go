package dal

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/yuxi311/webService/model/entity"
)

func AutoMigrate(db *gorm.DB) error {
	models := []interface{}{
		&entity.User{},
	}

	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			return fmt.Errorf("auto migrate: %T error: %v", model, err)
		}
	}
	return nil
}

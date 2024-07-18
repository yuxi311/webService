package dal

import (
	"fmt"

	"github.com/yuxi311/webService/model/entity"
)

func AutoMigrate() error {
	db := DB()
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

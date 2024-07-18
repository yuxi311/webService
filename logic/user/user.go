package user

import (
	// "strings"
	// "gorm.io/gorm"

	"github.com/yuxi311/webService/dal"
	"github.com/yuxi311/webService/model/entity"
)

func CreateNewUser(name string, username string, password string, role int32) error {
	db := dal.DB()

	user := &entity.User{
		Name:     name,
		Username: username,
		Password: password,
		Role:     role,
	}

	db.Create(&user)
	return nil
}

func QueryUsers() []entity.User {
	var users []entity.User
	db := dal.DB()

	db.Find(&users)
	return users
}

func QueryUser(id int) entity.User {
	var user entity.User
	db := dal.DB()

	db.Where("id = ?", id).Find(&user)

	return user
}

func DeleteUser(id int) error {
	db := dal.DB()

	db.Where("id = ?", id).Delete(&entity.User{})
	return nil
}

func UpdateUser(id int, name string, password string, role int32) error {
	db := dal.DB()

	db.Model(&entity.User{}).Where("id = ?", id).Updates(entity.User{Name: name, Password: password, Role: role})
	return nil
}

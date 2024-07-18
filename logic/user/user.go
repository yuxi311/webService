package user

import (
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

func QueryUser(id uint64) entity.User {
	var user entity.User
	db := dal.DB()

	db.Where("id = ?", id).Find(&user)

	return user
}

func DeleteUser(id uint64) error {
	db := dal.DB()

	db.Where("id = ?", id).Delete(&entity.User{})
	return nil
}

func UpdateUser(id uint64, name string, password string, role int32) error {
	db := dal.DB()

	db.Model(&entity.User{}).Where("id = ?", id).Updates(entity.User{Name: name, Password: password, Role: role})
	return nil
}

func QueryUserByUsername(username string) entity.User {
	var user entity.User
	db := dal.DB()

	db.Where("username = ?", username).Find(&user)

	return user
}

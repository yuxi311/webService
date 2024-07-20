package user

import (
	"context"
	"strconv"
	"time"

	"github.com/yuxi311/webService/dal"
	"github.com/yuxi311/webService/model/entity"
	"github.com/yuxi311/webService/pkg/logger"
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

	ctx := context.Background()
	rdb := dal.RDB()
	idString := strconv.Itoa(int(id))

	result, err := rdb.HGetAll(ctx, idString).Result()
	if len(result) == 0 || err != nil {
		db.Where("id = ?", id).Find(&user)

		values := structToMap(user)
		err = rdb.HMSet(ctx, idString, values).Err()
		if err != nil {
			logger.Errorf("set redis error: %v", err)
		}
	} else {
		user = mapToStruct(result)
	}
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

func structToMap(user entity.User) map[string]interface{} {
	layout := "2006-01-02T15:04:05.999999Z07:00"
	values := map[string]interface{}{
		"id":          user.ID,
		"name":        user.Name,
		"username":    user.Username,
		"password":    user.Password,
		"role":        user.Role,
		"description": user.Description,
		"create_at":   user.CreateAt.Format(layout),
		"update_at":   user.UpdateAt.Format(layout),
	}

	return values
}

func mapToStruct(result map[string]string) entity.User {
	layout := "2006-01-02T15:04:05.999999Z07:00"

	id, _ := strconv.ParseUint(result["id"], 10, 64)
	role, _ := strconv.ParseInt(result["role"], 10, 32)
	createAt, _ := time.Parse(layout, result["create_at"])
	updateAt, _ := time.Parse(layout, result["update_at"])

	user := entity.User{
		ID:          id,
		Name:        result["name"],
		Username:    result["username"],
		Password:    result["password"],
		Role:        int32(role),
		Description: result["description"],
		CreateAt:    createAt,
		UpdateAt:    updateAt,
	}

	return user
}

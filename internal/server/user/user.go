package user

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yuxi311/webService/logic/user"
	"github.com/yuxi311/webService/pkg/httpresponse"
	"github.com/yuxi311/webService/pkg/utils"
)

type NewUserBody struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        int32  `json:"role"`
	Description string `json:"description"`
}

type UpdateUserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int32  `json:"role"`
}

type UsersBody struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	Role        int32  `json:"role"`
	Description string `json:"description"`
}

type UserBody struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Role        int32     `json:"role"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}

func getAllUsersHandler(c *gin.Context) {
	var resp []UsersBody
	userList := user.QueryUsers()

	for _, userInfo := range userList {
		resp = append(resp, UsersBody{
			ID:          userInfo.ID,
			Name:        userInfo.Name,
			Username:    userInfo.Username,
			Role:        userInfo.Role,
			Description: userInfo.Description,
		})
	}

	httpresponse.Succeed(c, resp)
}

func getUserHandler(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		httpresponse.Fail(c, 10004, "tran userId failed")
		return
	}

	userInfo := user.QueryUser(uint64(userId))

	resp := UserBody{
		ID:          userInfo.ID,
		Name:        userInfo.Name,
		Username:    userInfo.Username,
		Password:    userInfo.Password,
		Role:        userInfo.Role,
		Description: userInfo.Description,
		CreateAt:    userInfo.CreateAt,
		UpdateAt:    userInfo.UpdateAt,
	}

	httpresponse.Succeed(c, resp)
}

func createUserHandler(c *gin.Context) {
	req := NewUserBody{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httpresponse.Fail(c, 10010, err.Error())
		return
	}

	hashedPassword := utils.EncodePassword(req.Password)
	if err := user.CreateNewUser(req.Name, req.Username, hashedPassword, req.Role); err != nil {
		httpresponse.Fail(c, 10006, "create new user failed")
		return
	}

	httpresponse.Succeed(c, nil)
}

func deleteUserHandler(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		httpresponse.Fail(c, 10004, "tran userId failed")
		return
	}

	user.DeleteUser(uint64(userId))
	httpresponse.Succeed(c, nil)
}

func updateUserHandler(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		httpresponse.Fail(c, 10004, "tran userId failed")
		return
	}

	req := UpdateUserBody{}
	if err := c.ShouldBindJSON(&req); err != nil {
		httpresponse.Fail(c, 10011, err.Error())
		return
	}

	hashedPassword := utils.EncodePassword(req.Password)
	if err := user.UpdateUser(uint64(userId), req.Username, hashedPassword, req.Role); err != nil {
		httpresponse.Fail(c, 10003, "update user failed")
		return
	}

	httpresponse.Succeed(c, nil)
}

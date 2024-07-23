package login

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/yuxi311/webService/logic/user"
	"github.com/yuxi311/webService/pkg/httpresponse"
	"github.com/yuxi311/webService/pkg/jwt"
	"github.com/yuxi311/webService/pkg/kafka"
	"github.com/yuxi311/webService/pkg/utils"
)

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRespBody struct {
	Token string `json:"token"`
	Role  int32  `json:"role"`
}

type LoginMessage struct {
	ConnectedAt time.Time `json:"connected_at"`
	Username    string    `json:"username"`
	Status      int       `json:"status"`
}

func loginHandler(c *gin.Context) {
	req := LoginBody{}

	if err := c.ShouldBindJSON(&req); err != nil {
		httpresponse.Fail(c, 10007, "parse request body error")
		return
	}

	hashedPassword := utils.EncodePassword(req.Password)
	user := user.QueryUserByUsername(req.Username)

	if user.Username == "" {
		httpresponse.Fail(c, 10001, "username not found")
		return
	}

	if hashedPassword != user.Password {
		httpresponse.Fail(c, 10002, "password error")
		return
	}

	token, err := jwt.CreateToken(user.Username, user.Role)
	if err != nil {
		httpresponse.Fail(c, 10008, err.Error())
		return
	}

	msg := LoginMessage{
		ConnectedAt: time.Now(),
		Username:    req.Username,
		Status:      1,
	}
	msgData, err := json.Marshal(msg)
	if err != nil {
		httpresponse.Fail(c, 10018, err.Error())
		return
	}
	err = kafka.ProduceMessage(msgData)
	if err != nil {
		httpresponse.Fail(c, 10015, err.Error())
		return
	}

	resp := LoginRespBody{
		Token: token,
		Role:  user.Role,
	}

	httpresponse.Succeed(c, resp)
}

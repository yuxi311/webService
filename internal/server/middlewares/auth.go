package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/yuxi311/webService/pkg/httpresponse"
	"github.com/yuxi311/webService/pkg/jwt"
)

func AuthToken(c *gin.Context) {
	req := c.Request

	authHeader := req.Header.Get("Authorization")

	if len(authHeader) == 0 {
		httpresponse.Fail(c, 10012, "missing token")
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	err := jwt.ParseToken(tokenString)
	if err != nil {
		httpresponse.Fail(c, 10013, err.Error())
		return
	}
}

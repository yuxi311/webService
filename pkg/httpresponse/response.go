package httpresponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Succeed(ctx *gin.Context, data ...interface{}) {
	if len(data) > 0 && data[0] != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, data[0])
	} else {
		ctx.Status(http.StatusOK)
	}
}

func Fail(ctx *gin.Context, code int, message string) {
	errMessage := ErrMessage{
		Code: code,
		Message: message,
	}
 	ctx.AbortWithStatusJSON(http.StatusInternalServerError, errMessage)
}

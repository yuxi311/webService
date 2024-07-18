package httpresponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Succeed(ctx *gin.Context, data ...interface{}) {
	if len(data) > 0 && data[0] != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, data[0])
	} else {
		ctx.Status(http.StatusOK)
	}
}

func Fail(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
}

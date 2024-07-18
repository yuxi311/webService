package login

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	loginRouter := router.Group("/login")
	loginRouter.POST("", loginHandler)
}

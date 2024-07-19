package login

import (
	"github.com/gin-gonic/gin"
)

func RegisterNoAuthRoutes(router *gin.RouterGroup) {
	loginRouter := router.Group("/login")
	loginRouter.POST("", loginHandler)
}

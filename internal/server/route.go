package server

import (
	"github.com/gin-gonic/gin"

	"github.com/yuxi311/webService/internal/server/user"
)

func setupRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api")
	user.RegisterRoutes(apiRouter)
}

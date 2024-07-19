package server

import (
	"github.com/gin-gonic/gin"

	"github.com/yuxi311/webService/internal/server/login"
	"github.com/yuxi311/webService/internal/server/middlewares"
	"github.com/yuxi311/webService/internal/server/user"
)

func setupRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api")
	apiRouter.Use(middlewares.AuthToken)

	user.RegisterRoutes(apiRouter)
}

func setupNoAuthRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api")
	login.RegisterNoAuthRoutes(apiRouter)
}

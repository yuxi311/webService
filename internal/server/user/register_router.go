package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	userRouter := router.Group("/users")
	userRouter.GET("", getAllUsersHandler)
	userRouter.GET("/:userId", getUserHandler)
	userRouter.POST("", createUserHandler)
	userRouter.DELETE("/:userId", deleteUserHandler)
	userRouter.PUT("/:userId", updateUserHandler)
}

package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Serve(port int) {
	gr := gin.New()
	gr.GET("/api/get", getapi)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: gr,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func getapi(c *gin.Context) {
	resp := make(map[string]string)
	resp["port"] = "90"
	succeed(c, resp)
}

// Succeed will response with success status
func succeed(ctx *gin.Context, data ...interface{}) {
	if len(data) > 0 && data[0] != nil {
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, data[0])
	} else {
		ctx.Status(http.StatusOK)
	}
}

package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Serve(port int) {
	gr := gin.New()
	setupRoutes(gr)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: gr,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

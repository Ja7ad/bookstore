package app

import (
	"bookstore/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default() // router for routing driving requests
)

func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

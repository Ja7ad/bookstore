package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default() // router for routing driving requests
)

func StartApplication() {
	mapUrls()
	router.Run(":8080")
}

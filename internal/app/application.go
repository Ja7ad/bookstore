package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default() // router for routing driving requests
)

func StartApplication() {
	mapUrls()
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

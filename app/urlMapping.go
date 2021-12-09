package app

import "bookstore/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}

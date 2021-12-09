package app

import (
	"bookstore/controllers/healthy"
	"bookstore/controllers/users"
)

func mapUrls() {
	router.GET("/ping", healthy.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}

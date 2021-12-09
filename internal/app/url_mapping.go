package app

import (
	"bookstore/internal/controllers/healthy"
	"bookstore/internal/controllers/users"
)

func mapUrls() {
	router.GET("/ping", healthy.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}

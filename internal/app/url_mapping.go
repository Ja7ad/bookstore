package app

import (
	"bookstore/internal/controllers/healthy"
	"bookstore/internal/controllers/users"
)

func mapUrls() {
	router.GET("/ping", healthy.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
}

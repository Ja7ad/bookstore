package app

import (
	"bookstore/pkg/oauth/clients/cassandra"
	"bookstore/pkg/oauth/domain"
	"bookstore/pkg/oauth/http"
	"bookstore/pkg/oauth/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	defer session.Close()

	atHandler := http.NewHandler(domain.NewTokenService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}

package app

import (
	"bookstore/internal/pkg/oauth/clients/cassandra"
	"bookstore/internal/pkg/oauth/domain"
	"bookstore/internal/pkg/oauth/http"
	"bookstore/internal/pkg/oauth/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session := cassandra.GetSession()
	defer session.Close()

	atHandler := http.NewHandler(domain.NewTokenService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}

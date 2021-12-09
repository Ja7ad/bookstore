package controllers

import "github.com/gin-gonic/gin"

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

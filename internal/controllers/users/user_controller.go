package users

import (
	"bookstore/errors/restError"
	"bookstore/internal/domain/users"
	"bookstore/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := restError.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusCreated, result)
	return
}

// GetUser returns a user by its id
func GetUser(ctx *gin.Context) {
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if userErr != nil {
		err := restError.NewBadRequestError("user id should be a number")
		ctx.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// UpdateUser updates a user by its id
func UpdateUser(ctx *gin.Context) {
	userId, userErr := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if userErr != nil {
		err := restError.NewBadRequestError("user id should be a number")
		ctx.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := restError.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	result, err := services.UpdateUser(user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, result)

}

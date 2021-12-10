package users

import (
	"bookstore/internal/domain/users"
	"bookstore/internal/services"
	"bookstore/pkg/errors/restError"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// getUserId validate user id from url parameter
func getUserId(userIdParam string) (int64, *restError.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, restError.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

// Create new user
func Create(ctx *gin.Context) {
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := restError.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}

	ctx.JSON(http.StatusCreated, result.Marshal(ctx.GetHeader("X-Public") == "true"))
	return
}

// Get returns a user by its id
func Get(ctx *gin.Context) {
	userId, idErr := getUserId(ctx.Param("user_id"))
	if idErr != nil {
		ctx.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		ctx.JSON(getErr.Status, getErr)
		return
	}

	ctx.JSON(http.StatusOK, user.Marshal(ctx.GetHeader("X-Public") == "true"))
}

// Update updates a user by its id
func Update(ctx *gin.Context) {
	userId, idErr := getUserId(ctx.Param("user_id"))
	if idErr != nil {
		ctx.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := restError.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	// isPartial check method is Patch
	isPartial := ctx.Request.Method == http.MethodPatch

	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, result)

}

// Delete delete user from database by id
func Delete(ctx *gin.Context) {
	userId, idErr := getUserId(ctx.Param("user_id"))
	if idErr != nil {
		ctx.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search user in database
func Search(ctx *gin.Context) {
	status := ctx.Query("status")

	usersList, err := services.UsersService.SearchUser(status)
	if err != nil {
		ctx.JSON(err.Status, err)
	}

	ctx.JSON(http.StatusOK, usersList.Marshal(ctx.GetHeader("X-Public") == "true"))
}

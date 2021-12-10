package http

import (
	"bookstore/internal/pkg/oauth/domain"
	"bookstore/pkg/errors/restError"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service domain.Service
}

func NewHandler(service domain.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(ctx *gin.Context) {
	accessToken, err := h.service.GetById(ctx.Param("access_token_id"))
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(ctx *gin.Context) {
	var at domain.AccessToken
	if err := ctx.ShouldBindJSON(&at); err != nil {
		restErr := restError.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	if err := h.service.Create(at); err != nil {
		ctx.JSON(err.Status, err)
		return
	}

	ctx.JSON(http.StatusCreated, at)
}

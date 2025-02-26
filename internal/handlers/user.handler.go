package handlers

import (
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/internal/usecases"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/helpers"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/models"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

type UserHandler struct {
	userUsecase usecases.UserUsecaseImpl
	i18nBundle  *i18n.Bundle
}

func NewUserHandler(userUsecase usecases.UserUsecaseImpl, i18nBundle *i18n.Bundle) *UserHandler {
	return &UserHandler{userUsecase: userUsecase, i18nBundle: i18nBundle}
}

func (r *UserHandler) Register(c *gin.Context) {
	var payload models.RegisterPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, helpers.NewBadRequestError("invalid payload", err).
			Applied(c, r.i18nBundle))
	}
	err := r.userUsecase.Register(c, &payload)
	if err != nil {
		c.JSON(err.Code, err.Applied(c, r.i18nBundle))
	}
	successResponse := helpers.NewSuccess("success.global.created", nil, nil, http.StatusCreated)
	c.JSON(http.StatusCreated, successResponse.Applied(c, r.i18nBundle))
}

package handlers

import (
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/internal/usecases"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/configs"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/constants"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/models"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/response"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

type UserHandler struct {
	userUsecase usecases.UserUsecaseImpl
	i18nBundle  *i18n.Bundle
	cfg         *configs.AppConfig
	response    *response.ResponseFormatter
}

func NewUserHandler(userUsecase usecases.UserUsecaseImpl, i18nBundle *i18n.Bundle, cfg *configs.AppConfig) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
		i18nBundle:  i18nBundle,
		cfg:         cfg,
		response:    response.NewFormatter(i18nBundle, nil, cfg.App.IsDevelopment),
	}
}

func (r *UserHandler) Register(c *gin.Context) {
	var payload models.RegisterPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, r.response.FormatError(c, response.NewAPIError(constants.ErrorBadRequestType, constants.ErrorPayload, err, nil)))
		return
	}
	err := r.userUsecase.Register(c, &payload)
	if err != nil {
		c.JSON(err.HTTPStatus(), r.response.FormatError(c, err))
		return
	}
	successResponse := response.NewSuccess(constants.SuccessCreated, nil, nil, http.StatusCreated)
	c.JSON(http.StatusCreated, r.response.FormatSuccess(c, successResponse))
	return
}

package handlers

import (
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/internal/usecases"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecases.UserUsecaseImpl
}

func NewUserHandler(userUsecase usecases.UserUsecaseImpl) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (r UserHandler) Register(c *gin.Context) error {
	var payload models.RegisterPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	err := r.userUsecase.Register(c, &payload)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "success"})
	return nil
}

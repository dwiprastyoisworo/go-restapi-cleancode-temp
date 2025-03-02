package usecases

import (
	"context"
	"errors"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/internal/repositories"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/constants"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/helpers"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/models"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/response"
	"gorm.io/gorm"
)

type UserUsecase struct {
	repo     repositories.RepositoryImpl[models.Users]
	repoUser repositories.UserRepositoryImpl
	db       *gorm.DB
}

func NewUserUsecase(repo repositories.RepositoryImpl[models.Users], repoUser repositories.UserRepositoryImpl, db *gorm.DB) UserUsecaseImpl {
	return &UserUsecase{repo: repo, repoUser: repoUser, db: db}
}

type UserUsecaseImpl interface {
	Register(ctx context.Context, payload *models.RegisterPayload) *response.APIError
}

func (u UserUsecase) Register(ctx context.Context, payload *models.RegisterPayload) *response.APIError {
	tx := u.db.Begin().WithContext(ctx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// get user by username
	err := u.checkUsernameExists(tx, payload.Username)
	if err != nil {
		return response.NewAPIError(constants.ErrorConflictType, constants.DataExists, err, nil)
	}

	// generate password hash
	hash, err := helpers.HashPassword(payload.Password)
	if err != nil {
		return response.NewAPIError(constants.ErrorConflictType, constants.ErrorUserHash, err, nil)
	}

	// create user
	userMap := map[string]interface{}{
		"username": payload.Username,
		"password": hash,
		"email":    payload.Email,
	}

	err = u.repo.Create(tx, userMap)
	if err != nil {
		return response.NewAPIError(constants.ErrorConflictType, constants.ErrorCreated, err, nil)
	}
	return nil

}

func (u UserUsecase) checkUsernameExists(tx *gorm.DB, username string) error {
	users, _ := u.repo.DynamicQuery(tx, map[string]string{"username": username})
	if len(users) > 0 {
		return errors.New("username already exists")
	}
	return nil
}

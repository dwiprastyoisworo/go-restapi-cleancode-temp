package usecases

import (
	"context"
	"errors"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/internal/repositories"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/helpers"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/models"
	"gorm.io/gorm"
	"log"
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
	Register(ctx context.Context, payload *models.RegisterPayload) error
}

func (u UserUsecase) Register(ctx context.Context, payload *models.RegisterPayload) error {
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
		log.Print(err)
		return err
	}

	// generate password hash
	hash, err := helpers.HashPassword(payload.Password)
	if err != nil {
		log.Print(err)
		return errors.New("failed to hash password")
	}

	// create user
	user := &models.Users{
		Username: payload.Username,
		Password: hash,
		Email:    payload.Email,
		FullName: payload.FullName,
	}

	err = u.repo.Create(tx, user)
	if err != nil {
		log.Print(err)
		return errors.New("failed to create user")
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

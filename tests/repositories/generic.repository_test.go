package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/internal/repositories"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	// Membuat GORM DB dengan mock
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	// Menyiapkan ekspektasi untuk query INSERT
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" \("email","fullName","password","username"\) VALUES \(\$1,\$2,\$3,\$4\) RETURNING "id"`).
		WithArgs("email@gmail.com", "Admin", "admin", "admin").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	// Membuat repository
	repo := repositories.NewRepository[models.Users]()
	err = repo.Create(gormDB, map[string]interface{}{
		"username": "admin",
		"password": "admin",
		"email":    "email@gmail.com",
		"fullName": "Admin",
	})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new record", err)
	}

	// Memastikan semua ekspektasi telah terpenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

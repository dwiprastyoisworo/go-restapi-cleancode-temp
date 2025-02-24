package repositories

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/internal/repositories"
	"github.com/dwiprastyoisworo/go-restapi-cleancode-temp/lib/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
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
	mock.ExpectExec("INSERT INTO \"users\" (.+) VALUES ($1,$2,$3,$4,$5,$6)").
		WithArgs("admin", "admin", "email@gmail.com", "Admin", sqlmock.AnyArg(), sqlmock.AnyArg()). // Gunakan sqlmock.AnyArg() untuk waktu
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Membuat repository
	repo := repositories.NewRepository[models.Users]()
	err = repo.Create(gormDB, &models.Users{
		Username:  "admin",
		Password:  "admin",
		Email:     "email@gmail.com",
		FullName:  "Admin",
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"), // Format time.Time ke string
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"), // Gunakan time.Time langsung
	})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when creating a new record", err)
	}

	// Memastikan semua ekspektasi telah terpenuhi
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

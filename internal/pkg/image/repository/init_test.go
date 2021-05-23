package repository

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func TestGetImages(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	uuid := uuid.New()
	userId := 1

	rows := sqlmock.NewRows([]string{"photoUuid"}).AddRow(uuid)

	mock.
		ExpectQuery("SELECT photoUuid FROM photos WHERE user").
		WithArgs(userId).
		WillReturnRows(rows)

	repo := PostgresRepository{
		Db: sqlx.NewDb(db, "psx"),
	}

	_, err = repo.GetImages(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

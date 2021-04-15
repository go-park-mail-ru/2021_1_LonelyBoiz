package repository

import (
	"database/sql/driver"
	"regexp"
	model "server/internal/pkg/models"
	"server/internal/pkg/user/repository"
	"testing"

	"github.com/jmoiron/sqlx"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.
		NewRows([]string{"id", "email", "name", "birthday", "instagram", "description", "city", "sex", "datepreference", "isactive", "isdeleted"})
	expect := model.User{
		Id:             1,
		Email:          "email",
		Name:           "Name",
		Birthday:       1,
		Instagram:      "inst",
		Description:    "desc",
		City:           "city",
		Sex:            "male",
		DatePreference: "male",
		IsActive:       true,
		IsDeleted:      false,
	}
	rows = rows.AddRow(
		expect.Id,
		expect.Email,
		expect.Name,
		expect.Birthday,
		expect.Instagram,
		expect.Description,
		expect.City,
		expect.Sex,
		expect.DatePreference,
		expect.IsActive,
		expect.IsDeleted,
	)

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT id, 
			email,
    		name,
    		birthday,
			instagram,
    		description,
    		city,
    		sex,
    		datepreference,
    		isactive,
    		isdeleted
		FROM users WHERE id = $1`)).
		WithArgs(driver.Value(1)).
		WillReturnRows(rows)

	repo := repository.UserRepository{DB: sqlx.NewDb(db, "pgx")}

	_, err = repo.GetUser(1)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

package repository

import (
	"reflect"
	"server/internal/pkg/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func TestCheckMail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	email := "example@mail.ru"

	rows := sqlmock.NewRows([]string{"email"}).AddRow(email)

	mock.
		ExpectQuery("SELECT email FROM users WHERE email").
		WithArgs(email).
		WillReturnRows(rows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.CheckMail(email)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, true) {
		t.Errorf("results not match, want %v, have %v", true, res)
		return
	}
}

func TestGetPassWithEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	email := "example@mail.ru"
	pass := []byte{1, 2, 3}

	rows := sqlmock.NewRows([]string{"passwordhash"}).AddRow(pass)

	mock.
		ExpectQuery("SELECT passwordHash FROM users WHERE email").
		WithArgs(email).
		WillReturnRows(rows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.GetPassWithEmail(email)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, pass) {
		t.Errorf("results not match, want %v, have %v", pass, res)
		return
	}
}

func TestGetPassWithId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	id := 1
	pass := []byte{1, 2, 3}

	rows := sqlmock.NewRows([]string{"passwordhash"}).AddRow(pass)

	mock.
		ExpectQuery("SELECT passwordHash FROM users WHERE id").
		WithArgs(id).
		WillReturnRows(rows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.GetPassWithId(id)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, pass) {
		t.Errorf("results not match, want %v, have %v", pass, res)
		return
	}
}

func TestSignIn(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	user := models.User{
		Id:             1,
		Email:          "exmaple@mail.ru",
		Name:           "name",
		Birthday:       123,
		Description:    "desc",
		City:           "city",
		Sex:            "male",
		Instagram:      "inst",
		PasswordHash:   []byte{1, 2},
		DatePreference: "male",
		IsActive:       true,
		IsDeleted:      true,
		Photos:         []uuid.UUID{},
	}

	rows := sqlmock.NewRows([]string{
		"id",
		"email",
		"name",
		"birthday",
		"description",
		"city",
		"sex",
		"instagram",
		"passwordhash",
		"datepreference",
		"isactive",
		"isdeleted"}).AddRow(
		user.Id,
		user.Email,
		user.Name,
		user.Birthday,
		user.Description,
		user.City,
		user.Sex,
		user.Instagram,
		user.PasswordHash,
		user.DatePreference,
		user.IsActive,
		user.IsDeleted,
	)

	rows2 := sqlmock.NewRows([]string{"photoid"}).AddRow(user.Photos[0])

	mock.
		ExpectQuery("SELECT id, email, name, birthday, description, city").
		WithArgs(user.Email).
		WillReturnRows(rows)

	mock.
		ExpectQuery("SELECT photoId FROM photos WHERE userid").
		WithArgs(user.Id).
		WillReturnRows(rows2)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.SignIn(user.Email)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, user) {
		t.Errorf("results not match, want %v, have %v", user, res)
		return
	}
}

func TestAddPhoto(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	photoid := 1

	rows := sqlmock.NewRows([]string{"photoid"}).AddRow(photoid)

	userid := 1
	value := "abc"

	mock.
		ExpectQuery("INSERT INTO ").
		WithArgs(userid, value).
		WillReturnRows(rows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.AddPhoto(userid, value)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, photoid) {
		t.Errorf("results not match, want %v, have %v", photoid, res)
		return
	}
}

func TestGetPhoto(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	photoid := 1
	value := "abc"

	rows := sqlmock.NewRows([]string{"value"}).AddRow(value)

	mock.
		ExpectQuery("SELECT value FROM photos").
		WithArgs(photoid).
		WillReturnRows(rows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.GetPhoto(photoid)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, value) {
		t.Errorf("results not match, want %v, have %v", value, res)
		return
	}
}

func TestGetChatById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	userid := 1

	chat := models.Chat{
		ChatId:              1,
		PartnerId:           1,
		PartnerName:         "Name",
		LastMessage:         "last message",
		LastMessageTime:     123,
		LastMessageAuthorId: 1,
		Photos:              []uuid.UUID{},
	}

	rows := sqlmock.NewRows([]string{
		"ChatId",
		"partnerId",
		"partnerName",
		"lastMessage",
		"lastMessageTime",
		"lastMessageAuthorId",
	}).AddRow(
		chat.ChatId,
		chat.PartnerId,
		chat.PartnerName,
		chat.LastMessage,
		chat.LastMessageTime,
		chat.LastMessageAuthorId,
	)

	rows2 := sqlmock.NewRows([]string{"photoid"}).AddRow(chat.Photos[0]).AddRow(chat.Photos[1])

	mock.
		ExpectQuery("SELECT chats.id AS chatId,").
		WithArgs(chat.ChatId, userid).
		WillReturnRows(rows)

	mock.
		ExpectQuery("SELECT photoId FROM photos").
		WithArgs(chat.PartnerId).
		WillReturnRows(rows2)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.GetChatById(chat.ChatId, userid)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, chat) {
		t.Errorf("results not match, want %v, have %v", chat, res)
		return
	}
}

func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	user := models.User{
		Id:             1,
		Email:          "exmaple@mail.ru",
		Name:           "name",
		Birthday:       123,
		Description:    "desc",
		City:           "city",
		Sex:            "male",
		Instagram:      "inst",
		PasswordHash:   []byte{1, 2},
		DatePreference: "male",
		IsActive:       true,
		IsDeleted:      true,
		Photos:         []uuid.UUID{},
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(user.Id)

	mock.
		ExpectQuery("INSERT INTO ").
		WithArgs(
			user.Email,
			user.Name,
			user.PasswordHash,
			user.Birthday,
			user.Description,
			user.City,
			user.Sex,
			user.DatePreference,
			user.IsActive,
			user.IsDeleted,
			user.Instagram,
		).
		WillReturnRows(rows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.AddUser(user)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, user.Id) {
		t.Errorf("results not match, want %v, have %v", user.Id, res)
		return
	}
}

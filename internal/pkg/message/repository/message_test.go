package repository

import (
	"reflect"
	"server/internal/pkg/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestCheckMail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	userId := 1
	name := "serega"

	rows := sqlmock.NewRows([]string{"name"}).AddRow(name)

	mock.
		ExpectQuery("SELECT name FROM user").
		WithArgs(userId).
		WillReturnRows(rows)

	repo := MessageRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.GetNameById(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, name) {
		t.Errorf("results not match, want %v, have %v", name, res)
		return
	}
}

func TestGetEmailById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	userId := 1
	email := "serega"

	rows := sqlmock.NewRows([]string{"email"}).AddRow(email)

	mock.
		ExpectQuery("SELECT email FROM use").
		WithArgs(userId).
		WillReturnRows(rows)

	repo := MessageRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.GetEmailById(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, email) {
		t.Errorf("results not match, want %v, have %v", email, res)
		return
	}
}

func TestCheckMessageForReacting(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	userId := 1
	messageId := 1
	authorId := 1

	rows := sqlmock.NewRows([]string{"authorId"}).AddRow(authorId)

	mock.
		ExpectQuery("SELECT authorId FROM messages").
		WithArgs(userId, messageId).
		WillReturnRows(rows)

	repo := MessageRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.CheckMessageForReacting(userId, messageId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, authorId) {
		t.Errorf("results not match, want %v, have %v", authorId, res)
		return
	}
}

func TestCheckChat(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	userId := 1
	chatId := 1

	rows := sqlmock.NewRows([]string{"id"}).AddRow(chatId)

	mock.
		ExpectQuery("SELECT id FROM chats").
		WithArgs(userId, chatId).
		WillReturnRows(rows)

	repo := MessageRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.CheckChat(userId, chatId)
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

func TestAddMessage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	newMessage := models.Message{
		AuthorId:     1,
		ChatId:       1,
		Text:         "text",
		Time:         123123123,
		Reaction:     -1,
		MessageId:    1,
		MessageOrder: 1,
	}
	rows := sqlmock.NewRows([]string{"messageid", "messageOrder"}).AddRow(newMessage.MessageId, newMessage.MessageOrder)

	mock.
		ExpectQuery("INSERT INTO messages").
		WithArgs(newMessage.ChatId,
			newMessage.AuthorId,
			newMessage.Text,
			sqlmock.AnyArg()).
		WillReturnRows(rows)

	repo := MessageRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	_, err = repo.AddMessage(newMessage.AuthorId,
		newMessage.ChatId,
		newMessage.Text)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

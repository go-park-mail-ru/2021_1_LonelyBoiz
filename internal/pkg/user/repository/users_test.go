package repository

import (
	"database/sql/driver"
	"reflect"
	"server/internal/pkg/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
		Id:               1,
		Email:            "exmaple@mail.ru",
		Name:             "name",
		Birthday:         123,
		Description:      "desc",
		City:             "city",
		Sex:              "male",
		Instagram:        "inst",
		PasswordHash:     []byte{1, 2},
		DatePreference:   "male",
		IsActive:         true,
		IsDeleted:        true,
		Photos:           []string{"1", "2"},
		Weight:           12,
		PartnerWeightTop: 12,
		PartnerWeightBot: 12,
		Height:           12,
		PartnerHeightTop: 12,
		PartnerHeightBot: 12,
		PartnerAgeTop:    12,
		PartnerAgeBot:    12,
		Interests:        make(pq.Int64Array, 0),
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
		"isdeleted",
		"photos",
		"weight",
		"partnerweighttop",
		"partnerweightbot",
		"height",
		"partnerheighttop",
		"partnerheightbot",
		"partneragetop",
		"partneragebot"}).AddRow(
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
		user.Photos,
		user.Weight,
		user.PartnerWeightBot,
		user.PartnerHeightTop,
		user.Height,
		user.PartnerHeightBot,
		user.PartnerWeightTop,
		user.PartnerAgeBot,
		user.PartnerAgeTop,
	)

	mock.
		ExpectQuery("SELECT id, email, name, birthday, description, city").
		WithArgs(user.Email).
		WillReturnRows(rows)

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
		Photos:              []string{"1", "2"},
	}

	rows := sqlmock.NewRows([]string{
		"chatid",
		"partnerid",
		"partnername",
		"lastmessage",
		"lastmessagetime",
		"lastmessageauthorid",
		"photos",
	}).AddRow(
		chat.ChatId,
		chat.PartnerId,
		chat.PartnerName,
		chat.LastMessage,
		chat.LastMessageTime,
		chat.LastMessageAuthorId,
		chat.Photos,
	)

	mock.
		ExpectQuery("SELECT chats.id AS chatid,").
		WithArgs(chat.ChatId, userid).
		WillReturnRows(rows)

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
		Photos:         []string{},
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
			user.Photos,
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

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	user := models.User{
		Id:               1,
		Email:            "exmaple@mail.ru",
		Name:             "name",
		Birthday:         123,
		Description:      "desc",
		City:             "city",
		Sex:              "male",
		Instagram:        "inst",
		PasswordHash:     []byte{1, 2},
		DatePreference:   "male",
		IsActive:         true,
		IsDeleted:        true,
		Photos:           []string{"1", "2"},
		Weight:           12,
		PartnerWeightTop: 12,
		PartnerWeightBot: 12,
		Height:           12,
		PartnerHeightTop: 12,
		PartnerHeightBot: 12,
		PartnerAgeTop:    12,
		PartnerAgeBot:    12,
		Interests:        make(pq.Int64Array, 0),
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
		"isdeleted",
		"photos",
		"weight",
		"partnerweighttop",
		"partnerweightbot",
		"height",
		"partnerheighttop",
		"partnerheightbot",
		"partneragetop",
		"partneragebot"}).AddRow(
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
		user.Photos,
		user.Weight,
		user.PartnerWeightBot,
		user.PartnerHeightTop,
		user.Height,
		user.PartnerHeightBot,
		user.PartnerWeightTop,
		user.PartnerAgeBot,
		user.PartnerAgeTop,
	)

	mock.
		ExpectQuery("SELECT id, email, name, birthday,").
		WithArgs(user.Id).
		WillReturnRows(rows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.GetUser(user.Id)
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

func TestCheckPermission(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	ownerId := 1
	getterId := 2

	rows := sqlmock.NewRows([]string{
		"ownerId",
	}).AddRow(
		ownerId,
	)

	mock.
		ExpectQuery("SELECT ownerId FROM secretPermiss").
		WithArgs(ownerId, getterId).
		WillReturnRows(rows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.CheckPermission(ownerId, getterId)
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

func TestReduceScrolls(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	scrolls := 1
	userId := 1

	rows := sqlmock.NewRows([]string{
		"scrolls",
	}).AddRow(
		scrolls,
	)

	mock.
		ExpectQuery("UPDATE users SET scrol").
		WithArgs(userId).
		WillReturnRows(rows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.ReduceScrolls(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, scrolls) {
		t.Errorf("results not match, want %v, have %v", scrolls, res)
		return
	}
}

func TestGetNewChatById(t *testing.T) {
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
		Photos:              []string{"1", "2"},
	}

	rows := sqlmock.NewRows([]string{
		"chatid",
		"partnerid",
		"partnername",
		"lastmessage",
		"lastmessagetime",
		"lastmessageauthorid",
		"photos",
	}).AddRow(
		chat.ChatId,
		chat.PartnerId,
		chat.PartnerName,
		chat.LastMessage,
		chat.LastMessageTime,
		chat.LastMessageAuthorId,
		chat.Photos,
	)

	mock.
		ExpectQuery("SELECT chats.id AS chatId,").
		WithArgs(chat.ChatId, userid).
		WillReturnRows(rows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.GetNewChatById(chat.ChatId, userid)
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

func TestCreateChat(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	chat := models.Chat{
		ChatId: 1,
	}

	userid1 := 1
	userid2 := 2

	rows := sqlmock.NewRows([]string{"id"}).AddRow(chat.ChatId)

	mock.
		ExpectQuery("INSERT INTO ").
		WithArgs(
			userid1,
			userid2,
		).
		WillReturnRows(rows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.CreateChat(userid1, userid2)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, chat.ChatId) {
		t.Errorf("results not match, want %v, have %v", chat.ChatId, res)
		return
	}
}

func TestUpdatePayment(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	amount := 10
	userid := 1

	mock.
		ExpectExec("UPDATE users").
		WithArgs(
			amount,
			userid,
		).WillReturnResult(driver.ResultNoRows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	err = repo.UpdatePayment(userid, amount)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestAddToSecreteAlbum(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	photos := []string{"1", "2"}
	userid := 1

	mock.
		ExpectExec("UPDATE secretPhotos").
		WithArgs(
			pq.Array(photos),
			userid,
		).WillReturnResult(driver.ResultNoRows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	err = repo.AddToSecreteAlbum(userid, photos)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestGetSecretePhotos(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	photos := []string{"1", "2"}
	userid := 1

	rows := sqlmock.NewRows([]string{"photos"}).AddRow(pq.Array(photos))

	mock.
		ExpectQuery("SELECT photos FROM secretPhotos").
		WithArgs(userid).WillReturnRows(rows)

	repo := UserRepository{
		DB: sqlx.NewDb(db, "psx"),
	}

	res, err := repo.GetSecretePhotos(userid)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(res, photos) {
		t.Errorf("results not match, want %v, have %v", photos, res)
		return
	}
}

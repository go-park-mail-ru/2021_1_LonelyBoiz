package usecase

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"server/internal/pkg/models"
	model "server/internal/pkg/models"
	mocks "server/internal/pkg/user/repository/mocks"
	user_proto "server/internal/user_server/delivery/proto"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const bcryptPass = "$2a$05$xaL9iW4Opbrgn52nyBO0/OIbfx1jjuIVy.SYCBG2VIqHzHnkj05le"

func TestUserUsecaseSignInNonValidEmail(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	user := models.User{
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
		Birthday:       0,
		Name:           "nick",
		DatePreference: "male",
		Sex:            "male"}

	_, code, err := UserUsecaseTest.SignIn(user)

	assert.Equal(t, 400, code)
	assert.NotEqual(t, nil, err)
}

func TestUserUsecaseSignInWrongPassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	user := models.User{
		Email:          "windes@ya.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		Birthday:       0,
		Name:           "nick",
		DatePreference: "male",
		Sex:            "male"}

	dbMock.EXPECT().GetPassWithEmail(user.Email).Return(pass, nil)

	_, code, err := UserUsecaseTest.SignIn(user)

	assert.Equal(t, 401, code)
	assert.NotEqual(t, nil, err)
}

func TestUserUsecaseSignInCheckPasswordError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	user := models.User{
		Email:          "windes@ya.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		Birthday:       0,
		Name:           "nick",
		DatePreference: "male",
		Sex:            "male"}

	dbMock.EXPECT().GetPassWithEmail(user.Email).Return(nil, errors.New("Some error"))

	_, code, err := UserUsecaseTest.SignIn(user)

	assert.Equal(t, 500, code)
	assert.Equal(t, nil, err)
}

func TestUserUsecaseSignInSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: &logrus.Entry{}},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	user1 := models.User{
		Id:             1,
		Email:          "windes@ya.ru",
		Password:       "12345678",
		SecondPassword: "12345678",
		PasswordHash:   nil,
		Name:           "nick",
		Birthday:       0,
		Description:    "",
		City:           "",
		Instagram:      "",
		Sex:            "male",
		DatePreference: "male",
		IsDeleted:      false,
		IsActive:       false,
		CaptchaToken:   "",
		Photos:         make([]string, 0),
	}

	user := models.User{
		Email:          "windes@ya.ru",
		Password:       "12345678",
		SecondPassword: "12345678",
		Birthday:       0,
		Name:           "nick",
		DatePreference: "male",
		Sex:            "male"}

	dbMock.EXPECT().SignIn(user.Email).Return(user1, nil)
	dbMock.EXPECT().GetPassWithEmail(user.Email).Return(pass, nil)

	returnUser, code, err := UserUsecaseTest.SignIn(user)

	assert.Equal(t, 200, code)
	assert.Equal(t, nil, err)
	assert.Equal(t, user1, returnUser)
}

func TestUserUsecaseSignInError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	user := models.User{
		Email:          "windes@ya.ru",
		Password:       "12345678",
		SecondPassword: "12345678",
		Birthday:       0,
		Name:           "nick",
		DatePreference: "male",
		Sex:            "male"}

	dbMock.EXPECT().SignIn(user.Email).Return(user, errors.New("Some error"))
	dbMock.EXPECT().GetPassWithEmail(user.Email).Return(pass, nil)

	_, code, err := UserUsecaseTest.SignIn(user)

	assert.Equal(t, 500, code)
	assert.Equal(t, nil, err)
}

func TestCreateChatNon_ReduceScrolls_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	like := models.Like{
		UserId:   1,
		Reaction: "asdf",
	}

	bufErr := errors.New("Some error")

	dbMock.EXPECT().ReduceScrolls(1).Return(1, bufErr)

	_, code, err := UserUsecaseTest.CreateChat(1, like)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestCreateChatNon_ReduceScrolls_Less0(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	like := models.Like{
		UserId:   1,
		Reaction: "asdf",
	}

	dbMock.EXPECT().ReduceScrolls(1).Return(-1, nil)

	_, code, err := UserUsecaseTest.CreateChat(1, like)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 402, code)
}

func TestCreateChatNonValidReaction(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	like := models.Like{
		UserId:   1,
		Reaction: "asdf",
	}

	dbMock.EXPECT().ReduceScrolls(1).Return(1, nil)

	_, code, err := UserUsecaseTest.CreateChat(1, like)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 400, code)
}

func TestCreateChatPermissionDenied(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	like := models.Like{
		UserId:   1,
		Reaction: "like",
	}

	dbMock.EXPECT().ReduceScrolls(1).Return(1, nil)
	dbMock.EXPECT().Rating(1, 1, "like").Return(int64(-1), nil)

	_, code, err := UserUsecaseTest.CreateChat(1, like)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 403, code)
}

func TestCreateChatRatingDatabaseError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	like := models.Like{
		UserId:   1,
		Reaction: "like",
	}

	dbMock.EXPECT().ReduceScrolls(1).Return(1, nil)
	dbMock.EXPECT().Rating(1, 1, "like").Return(int64(-1), errors.New("Some error"))

	_, code, err := UserUsecaseTest.CreateChat(1, like)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestCreateChatCheckReciprocityError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	like := models.Like{
		UserId:   1,
		Reaction: "like",
	}

	dbMock.EXPECT().ReduceScrolls(1).Return(1, nil)
	dbMock.EXPECT().Rating(1, 1, "like").Return(int64(1), nil)
	dbMock.EXPECT().CheckReciprocity(1, 1).Return(false, errors.New("Some error"))

	_, code, err := UserUsecaseTest.CreateChat(1, like)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestCreateChatCheckReciprocityFalse(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	like := models.Like{
		UserId:   1,
		Reaction: "like",
	}

	dbMock.EXPECT().ReduceScrolls(1).Return(1, nil)
	dbMock.EXPECT().Rating(1, 1, "like").Return(int64(1), nil)
	dbMock.EXPECT().CheckReciprocity(1, 1).Return(false, nil)

	_, code, err := UserUsecaseTest.CreateChat(1, like)
	assert.Equal(t, nil, err)
	assert.Equal(t, 204, code)
}

func TestCreateChatCreateChatError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	like := models.Like{
		UserId:   1,
		Reaction: "like",
	}

	dbMock.EXPECT().ReduceScrolls(1).Return(1, nil)
	dbMock.EXPECT().Rating(1, 1, "like").Return(int64(1), nil)
	dbMock.EXPECT().CheckReciprocity(1, 1).Return(true, nil)
	dbMock.EXPECT().CreateChat(1, 1).Return(-1, errors.New("Some error"))

	_, code, err := UserUsecaseTest.CreateChat(1, like)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestCreateChatGetNewChatByIdError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	like := models.Like{
		UserId:   1,
		Reaction: "like",
	}

	chat := models.Chat{
		ChatId:    1,
		PartnerId: 2,
	}

	dbMock.EXPECT().ReduceScrolls(1).Return(1, nil)
	dbMock.EXPECT().Rating(1, 1, "like").Return(int64(1), nil)
	dbMock.EXPECT().CheckReciprocity(1, 1).Return(true, nil)
	dbMock.EXPECT().CreateChat(1, 1).Return(1, nil)
	dbMock.EXPECT().GetNewChatById(1, 1).Return(chat, errors.New("Some error"))

	_, code, err := UserUsecaseTest.CreateChat(1, like)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestCreateChatSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	like := models.Like{
		UserId:   1,
		Reaction: "like",
	}

	chat := models.Chat{
		ChatId:    1,
		PartnerId: 2,
	}

	dbMock.EXPECT().ReduceScrolls(1).Return(1, nil)
	dbMock.EXPECT().Rating(1, 1, "like").Return(int64(1), nil)
	dbMock.EXPECT().CheckReciprocity(1, 1).Return(true, nil)
	dbMock.EXPECT().CreateChat(1, 1).Return(1, nil)
	dbMock.EXPECT().GetNewChatById(1, 1).Return(chat, nil)

	_, code, err := UserUsecaseTest.CreateChat(1, like)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, code)
}

func TestCreateFeedGetFeedDbError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	limit := 10
	userId := 1

	dbMock.EXPECT().GetFeed(userId, limit).Return(nil, errors.New("Some error"))

	_, code, err := UserUsecaseTest.CreateFeed(userId, limit)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestCreateFeedCreateFeedDbError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	limit := 10
	userId := 1

	dbMock.EXPECT().GetFeed(userId, limit).Return([]int{1, 2}, nil)
	dbMock.EXPECT().CreateFeed(userId).Return(errors.New("Some error"))

	_, code, err := UserUsecaseTest.CreateFeed(userId, limit)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestCreateFeedSecondGetFeedDbError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	limit := 10
	userId := 1

	dbMock.EXPECT().GetFeed(userId, limit).Return([]int{1, 2}, nil)
	dbMock.EXPECT().CreateFeed(userId).Return(nil)
	dbMock.EXPECT().GetFeed(userId, limit).Return(nil, errors.New("Some error"))

	_, code, err := UserUsecaseTest.CreateFeed(userId, limit)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestCreateFeedClearFeedDbError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	limit := 10
	userId := 1

	dbMock.EXPECT().GetFeed(userId, limit).Return([]int{1, 2}, nil)
	dbMock.EXPECT().CreateFeed(userId).Return(nil)
	dbMock.EXPECT().GetFeed(userId, limit).Return(nil, nil)
	dbMock.EXPECT().ClearFeed(userId).Return(errors.New("Some error"))

	_, code, err := UserUsecaseTest.CreateFeed(userId, limit)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestCreateFeedThirdGetFeedDbError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	limit := 10
	userId := 1

	dbMock.EXPECT().GetFeed(userId, limit).Return([]int{1, 2}, nil)
	dbMock.EXPECT().CreateFeed(userId).Return(nil)
	dbMock.EXPECT().GetFeed(userId, limit).Return(nil, nil)
	dbMock.EXPECT().ClearFeed(userId).Return(nil)
	dbMock.EXPECT().GetFeed(userId, limit).Return(nil, errors.New("Some error"))

	_, code, err := UserUsecaseTest.CreateFeed(userId, limit)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestCreateFeedSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	limit := 10
	userId := 1

	dbMock.EXPECT().GetFeed(userId, limit).Return([]int{1, 2}, nil)
	dbMock.EXPECT().CreateFeed(userId).Return(nil)
	dbMock.EXPECT().GetFeed(userId, limit).Return(nil, nil)
	dbMock.EXPECT().ClearFeed(userId).Return(nil)
	dbMock.EXPECT().GetFeed(userId, limit).Return(nil, nil)

	_, code, err := UserUsecaseTest.CreateFeed(userId, limit)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, code)
}

func TestChangeUserInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	newUser := models.User{
		Id:             1,
		Email:          "windes@ya.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	oldUser := models.User{
		Id:             1,
		Email:          "windes1@ya.ru",
		PasswordHash:   pass,
		Name:           "notnick",
		Birthday:       0,
		Description:    "",
		City:           "",
		Instagram:      "",
		Sex:            "female",
		DatePreference: "female",
		IsDeleted:      false,
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(pass, nil)
	dbMock.EXPECT().CheckMail(newUser.Email).Return(false, nil)
	dbMock.EXPECT().GetPhotos(newUser.Id).Return([]string{"1"}, nil)
	dbMock.EXPECT().ChangePassword(newUser.Id, gomock.Any()).Return(nil)
	dbMock.EXPECT().GetUser(newUser.Id).Return(oldUser, nil)
	dbMock.EXPECT().ChangeUser(gomock.Any()).Return(nil)
	dbMock.EXPECT().CleanFeed(newUser.Id).Return(nil)

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, code)
}

func TestChangeUserInfo_CLeanFeed_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	newUser := models.User{
		Id:             1,
		Email:          "windes@ya.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	oldUser := models.User{
		Id:             1,
		Email:          "windes1@ya.ru",
		PasswordHash:   pass,
		Name:           "notnick",
		Birthday:       0,
		Description:    "",
		City:           "",
		Instagram:      "",
		Sex:            "female",
		DatePreference: "female",
		IsDeleted:      false,
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(pass, nil)
	dbMock.EXPECT().CheckMail(newUser.Email).Return(false, nil)
	dbMock.EXPECT().GetPhotos(newUser.Id).Return([]string{"1"}, nil)
	dbMock.EXPECT().ChangePassword(newUser.Id, gomock.Any()).Return(nil)
	dbMock.EXPECT().GetUser(newUser.Id).Return(oldUser, nil)
	dbMock.EXPECT().ChangeUser(gomock.Any()).Return(nil)
	dbMock.EXPECT().CleanFeed(newUser.Id).Return(errors.New("Some error"))

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestChangeUserInfoPasswordValidationError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	newUser := models.User{
		Id:             1,
		Email:          "windes@ya.ru",
		Password:       "123",
		SecondPassword: "123",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 400, code)
}

func TestChangeUserInfoWrongSecondPassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	newUser := models.User{
		Id:             1,
		Email:          "windes@ya.ru",
		Password:       "12345678",
		SecondPassword: "123456789",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 400, code)
}

func TestChangeUserInfoGetPassError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	newUser := models.User{
		Id:             1,
		Email:          "windes@ya.ru",
		Password:       "12345678",
		SecondPassword: "12345678",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(nil, errors.New("Some error"))

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 400, code)
}

func TestChangeUserInfoWrongPassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	newUser := models.User{
		Id:             1,
		Email:          "windes@ya.ru",
		Password:       "12345678",
		SecondPassword: "12345678",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().GetPassWithId(newUser.Id).Return([]byte{1, 2}, nil)

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 400, code)
}

func TestChangeUserInfoNilPasswordError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	newUser := models.User{
		Id:             1,
		Email:          "windes@ya.ru",
		Password:       "12345678",
		SecondPassword: "12345678",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(nil, nil)

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 400, code)
}

func TestChangeUserInfoChangePasswordError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	newUser := models.User{
		Id:             1,
		Email:          "windes@ya.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().ChangePassword(newUser.Id, gomock.Any()).Return(errors.New("Some error"))
	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(pass, nil)

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 400, code)
}

func TestChangeUserInfoGetUserInfoError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	newUser := models.User{
		Id:             1,
		Email:          "windes@ya.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().ChangePassword(newUser.Id, gomock.Any()).Return(nil)
	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(pass, nil)
	dbMock.EXPECT().GetUser(newUser.Id).Return(models.User{}, errors.New("Some error"))

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestChangeUserInfoNonValidEmail(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	newUser := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "123456789",
		SecondPassword: "123456789",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().ChangePassword(newUser.Id, gomock.Any()).Return(nil)
	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(pass, nil)
	dbMock.EXPECT().GetUser(newUser.Id).Return(newUser, nil)

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 400, code)
}

func TestChangeUserInfoEmailIsSignedUpError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	newUser := models.User{
		Id:             1,
		Email:          "windes@mail.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().ChangePassword(newUser.Id, gomock.Any()).Return(nil)
	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(pass, nil)
	dbMock.EXPECT().GetUser(newUser.Id).Return(newUser, nil)
	dbMock.EXPECT().CheckMail(newUser.Email).Return(false, errors.New("Some error"))

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestChangeUserInfoEmailIsSignedUp(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	newUser := models.User{
		Id:             1,
		Email:          "windes@mail.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().ChangePassword(newUser.Id, gomock.Any()).Return(nil)
	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(pass, nil)
	dbMock.EXPECT().GetUser(newUser.Id).Return(newUser, nil)
	dbMock.EXPECT().CheckMail(newUser.Email).Return(true, nil)

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 400, code)
}

func TestChangeUserInfoNonValidSex(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	newUser := models.User{
		Id:             1,
		Email:          "windes@mail.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male1",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().ChangePassword(newUser.Id, gomock.Any()).Return(nil)
	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(pass, nil)
	dbMock.EXPECT().GetUser(newUser.Id).Return(newUser, nil)
	dbMock.EXPECT().CheckMail(newUser.Email).Return(false, nil)

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 400, code)
}

func TestChangeUserInfoNonValidPreference(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	newUser := models.User{
		Id:             1,
		Email:          "windes@mail.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male1",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().ChangePassword(newUser.Id, gomock.Any()).Return(nil)
	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(pass, nil)
	dbMock.EXPECT().GetUser(newUser.Id).Return(newUser, nil)
	dbMock.EXPECT().CheckMail(newUser.Email).Return(false, nil)

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 400, code)
}

func TestChangeUserInfoIsActiveError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	newUser := models.User{
		Id:             1,
		Email:          "windes@mail.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().ChangePassword(newUser.Id, gomock.Any()).Return(nil)
	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(pass, nil)
	dbMock.EXPECT().GetUser(newUser.Id).Return(newUser, nil)
	dbMock.EXPECT().CheckMail(newUser.Email).Return(true, nil)

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 400, code)
}

func TestChangeUserInfoIsActiveFalse(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	newUser := models.User{
		Id:             1,
		Email:          "windes@mail.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().ChangePassword(newUser.Id, gomock.Any()).Return(nil)
	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(pass, nil)
	dbMock.EXPECT().GetUser(newUser.Id).Return(newUser, nil)
	dbMock.EXPECT().CheckMail(newUser.Email).Return(false, nil)
	dbMock.EXPECT().GetPhotos(newUser.Id).Return(nil, nil)
	dbMock.EXPECT().ChangeUser(gomock.Any()).Return(nil)
	dbMock.EXPECT().CleanFeed(newUser.Id).Return(nil)

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, code)
}

func TestChangeUserInfoChangeUserError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass := []byte(bcryptPass)

	newUser := models.User{
		Id:             1,
		Email:          "windes@mail.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		OldPassword:    "12345678",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().ChangePassword(newUser.Id, gomock.Any()).Return(nil)
	dbMock.EXPECT().GetPassWithId(newUser.Id).Return(pass, nil)
	dbMock.EXPECT().GetUser(newUser.Id).Return(newUser, nil)
	dbMock.EXPECT().CheckMail(newUser.Email).Return(false, nil)
	dbMock.EXPECT().GetPhotos(newUser.Id).Return(nil, nil)
	dbMock.EXPECT().ChangeUser(gomock.Any()).Return(errors.New("Some error"))

	_, code, err := UserUsecaseTest.ChangeUserInfo(newUser, newUser.Id)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestValidateSignUpDataSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	newUser := models.User{
		Id:             1,
		Email:          "windes@mail.ru",
		Password:       "123456789",
		SecondPassword: "123456789",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	err := UserUsecaseTest.ValidateSignUpData(newUser)
	assert.Equal(t, nil, err)
}

func TestValidateSignUpDataError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	newUser := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "123456789",
		SecondPassword: "123456789",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	err := UserUsecaseTest.ValidateSignUpData(newUser)
	assert.NotEqual(t, nil, err)
}

func TestValidateSignUpDataPassError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	newUser := models.User{
		Id:             1,
		Email:          "windes@mail.ru",
		Password:       "123456789",
		SecondPassword: "1234567891",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	err := UserUsecaseTest.ValidateSignUpData(newUser)
	assert.NotEqual(t, nil, err)
}

func TestAddNewUserSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	newUser := models.User{
		Id:             1,
		Email:          "windes@mail.ru",
		Password:       "123456789",
		SecondPassword: "1234567891",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().CreateSecretAlbum(newUser.Id).Return(nil)
	dbMock.EXPECT().AddUser(gomock.Any()).Return(newUser.Id, nil)

	err := UserUsecaseTest.AddNewUser(&newUser)
	assert.Equal(t, nil, err)
}

func TestAddNewUserAddUserError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	newUser := models.User{
		Id:             1,
		Email:          "windes@mail.ru",
		Password:       "123456789",
		SecondPassword: "1234567891",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         make([]string, 0),
	}

	dbMock.EXPECT().AddUser(gomock.Any()).Return(newUser.Id, errors.New("Some error"))

	err := UserUsecaseTest.AddNewUser(&newUser)
	assert.NotEqual(t, nil, err)
}

func TestGetUserInfoById(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	newUser := models.User{
		Id:             1,
		Email:          "windes@mail.ru",
		Password:       "123456789",
		SecondPassword: "1234567891",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         nil,
	}

	dbMock.EXPECT().GetUser(1).Return(newUser, nil)

	_, err := UserUsecaseTest.GetUserInfoById(newUser.Id)
	assert.Equal(t, nil, err)
}

func TestParseJsonToUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	newUser := models.User{
		Id:             1,
		Email:          "windes@mail.ru",
		Password:       "123456789",
		SecondPassword: "1234567891",
		Name:           "nick",
		Birthday:       123123,
		Description:    "desc",
		City:           "city",
		Instagram:      "inst",
		Sex:            "male",
		DatePreference: "male",
		Photos:         nil,
	}

	buf, err := json.Marshal(newUser)
	assert.Equal(t, nil, err)

	r := ioutil.NopCloser(strings.NewReader(string(buf)))

	_, err = UserUsecaseTest.ParseJsonToUser(r)
	assert.Equal(t, nil, err)
}

func TestGetSecreteAlbum(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	photos := []string{"1", "2"}
	ownerId := 1
	getterId := 2

	dbMock.EXPECT().CheckPermission(ownerId, getterId).Return(true, nil)
	dbMock.EXPECT().GetSecretePhotos(ownerId).Return(photos, nil)

	res, code, err := UserUsecaseTest.GetSecreteAlbum(ownerId, getterId)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, code)
	assert.Equal(t, res, photos)
}

func TestGetSecreteAlbum_CheckPermission_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	ownerId := 1
	getterId := 2

	bufErr := errors.New("Some error")
	dbMock.EXPECT().CheckPermission(ownerId, getterId).Return(false, bufErr)

	res, code, err := UserUsecaseTest.GetSecreteAlbum(ownerId, getterId)
	assert.Equal(t, bufErr, err)
	assert.Equal(t, 500, code)
	assert.Equal(t, []string{}, res)
}

func TestGetSecreteAlbum_CheckPermission_NotOk(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	ownerId := 1
	getterId := 2

	dbMock.EXPECT().CheckPermission(ownerId, getterId).Return(false, nil)

	res, code, err := UserUsecaseTest.GetSecreteAlbum(ownerId, getterId)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 403, code)
	assert.Equal(t, []string{}, res)
}

func TestGetSecreteAlbum_GetSecretePhotos_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	bufErr := errors.New("Some error")
	photos := []string{"1", "2"}
	ownerId := 1
	getterId := 2

	dbMock.EXPECT().CheckPermission(ownerId, getterId).Return(true, nil)
	dbMock.EXPECT().GetSecretePhotos(ownerId).Return(photos, bufErr)

	res, code, err := UserUsecaseTest.GetSecreteAlbum(ownerId, getterId)
	assert.Equal(t, bufErr, err)
	assert.Equal(t, 500, code)
	assert.Equal(t, res, []string{})
}

func TestUnblockSecreteAlbum(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	ownerId := 1
	getterId := 2

	dbMock.EXPECT().UnblockSecreteAlbum(ownerId, getterId).Return(nil)

	code, err := UserUsecaseTest.UnblockSecreteAlbum(ownerId, getterId)
	assert.Equal(t, nil, err)
	assert.Equal(t, 204, code)
}

func TestUnblockSecreteAlbum_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	bufErr := errors.New("Some error")
	ownerId := 1
	getterId := 2

	dbMock.EXPECT().UnblockSecreteAlbum(ownerId, getterId).Return(bufErr)

	code, err := UserUsecaseTest.UnblockSecreteAlbum(ownerId, getterId)
	assert.Equal(t, bufErr, err)
	assert.Equal(t, 500, code)
}

func TestAddToSecreteAlbum(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	ownerId := 1
	photos := []string{"1", "2"}

	dbMock.EXPECT().AddToSecreteAlbum(ownerId, photos).Return(nil)

	code, err := UserUsecaseTest.AddToSecreteAlbum(ownerId, photos)
	assert.Equal(t, nil, err)
	assert.Equal(t, 204, code)
}

func TestAddToSecreteAlbum_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecaseTest := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: logrus.New().WithField("test", "test")},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	bufErr := errors.New("Some error")
	ownerId := 1
	photos := []string{"1", "2"}

	dbMock.EXPECT().AddToSecreteAlbum(ownerId, photos).Return(bufErr)

	code, err := UserUsecaseTest.AddToSecreteAlbum(ownerId, photos)
	assert.Equal(t, bufErr, err)
	assert.Equal(t, 500, code)
}

func TestProtoUser2User(t *testing.T) {
	user := model.User{
		Id:             1,
		Email:          "email",
		Password:       "pass",
		SecondPassword: "pass",
		PasswordHash:   nil,
		OldPassword:    "pass",
		Name:           "Serega",
		Birthday:       123,
		Description:    "desc",
		City:           "Moscow",
		Instagram:      "Inst",
		Sex:            "male",
		DatePreference: "female",
		IsDeleted:      false,
		IsActive:       true,
		Photos:         make([]string, 0),
		Interests:      make([]int64, 0),
	}

	protoUser := user_proto.User{
		Id:             1,
		Email:          "email",
		Password:       "pass",
		SecondPassword: "pass",
		PasswordHash:   nil,
		OldPassword:    "pass",
		Name:           "Serega",
		Birthday:       123,
		Description:    "desc",
		City:           "Moscow",
		Instagram:      "Inst",
		Sex:            "male",
		DatePreference: "female",
		IsDeleted:      false,
		IsActive:       true,
		Photos:         make([]string, 0),
		Interests:      make([]int64, 0),
	}

	UserUsecaseTest := UserUsecase{}

	res := UserUsecaseTest.ProtoUser2User(&protoUser)

	assert.Equal(t, res, user)
}

func TestUser2ProtoUser(t *testing.T) {
	user := model.User{
		Id:             1,
		Email:          "email",
		Password:       "pass",
		SecondPassword: "pass",
		PasswordHash:   nil,
		OldPassword:    "pass",
		Name:           "Serega",
		Birthday:       123,
		Description:    "desc",
		City:           "Moscow",
		Instagram:      "Inst",
		Sex:            "male",
		DatePreference: "female",
		IsDeleted:      false,
		IsActive:       true,
		Photos:         nil,
		Interests:      nil,
	}

	protoUser := user_proto.User{
		Id:             1,
		Email:          "email",
		Password:       "pass",
		SecondPassword: "pass",
		PasswordHash:   nil,
		OldPassword:    "pass",
		Name:           "Serega",
		Birthday:       123,
		Description:    "desc",
		City:           "Moscow",
		Instagram:      "Inst",
		Sex:            "male",
		DatePreference: "female",
		IsDeleted:      false,
		IsActive:       true,
		Photos:         nil,
		Interests:      nil,
	}

	UserUsecaseTest := UserUsecase{}

	res := UserUsecaseTest.User2ProtoUser(user)

	assert.Equal(t, res, &protoUser)
}

func TestSetCookie(t *testing.T) {
	token := "some string"
	res := http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().AddDate(0, 0, 1),
		SameSite: http.SameSiteStrictMode,
		Domain:   model.GetDomain(),
		Secure:   model.GetSecure(),
		HttpOnly: true,
		Path:     "/",
	}

	UserUsecaseTest := UserUsecase{}

	ret := UserUsecaseTest.SetCookie(token)

	assert.Equal(t, res.Secure, ret.Secure)
	assert.Equal(t, res.Domain, ret.Domain)
	assert.Equal(t, res.SameSite, ret.SameSite)
	assert.Equal(t, res.Value, ret.Value)
	assert.Equal(t, res.HttpOnly, ret.HttpOnly)
}

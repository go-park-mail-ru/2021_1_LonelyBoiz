package usecase

import (
	"errors"
	"fmt"
	"server/internal/pkg/models"
	"server/internal/pkg/user/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

//"server/internal/pkg/user/usecase"

func TestUserUsecaseSignInNonValidEmail(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	//logger1 := models.Logger{Logger: &logrus.Entry{}}

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

	pass, err := UserUsecaseTest.HashPassword("12345678")
	if err != nil {
		fmt.Println("bcrypt error:", err)
	}

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

	pass, err := UserUsecaseTest.HashPassword("12345678")
	if err != nil {
		fmt.Println("bcrypt error:", err)
	}

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
		Photos:         make([]int, 0),
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

	pass, err := UserUsecaseTest.HashPassword("12345678")
	if err != nil {
		fmt.Println("bcrypt error:", err)
	}

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

	dbMock.EXPECT().Rating(1, 1, "like").Return(int64(1), nil)
	dbMock.EXPECT().CheckReciprocity(1, 1).Return(true, nil)
	dbMock.EXPECT().CreateChat(1, 1).Return(-1, errors.New("Some error"))

	_, code, err := UserUsecaseTest.CreateChat(1, like)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, code)
}

func TestCreateChatGetPhotosError(t *testing.T) {
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

	dbMock.EXPECT().Rating(1, 1, "like").Return(int64(1), nil)
	dbMock.EXPECT().CheckReciprocity(1, 1).Return(true, nil)
	dbMock.EXPECT().CreateChat(1, 1).Return(1, nil)
	dbMock.EXPECT().GetPhotos(1).Return(nil, errors.New("Some error"))

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

	dbMock.EXPECT().Rating(1, 1, "like").Return(int64(1), nil)
	dbMock.EXPECT().CheckReciprocity(1, 1).Return(true, nil)
	dbMock.EXPECT().CreateChat(1, 1).Return(1, nil)
	dbMock.EXPECT().GetPhotos(1).Return([]int{1, 2}, nil)

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

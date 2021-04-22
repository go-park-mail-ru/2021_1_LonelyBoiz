package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"server/internal/pkg/models"
	"server/internal/pkg/user/repository/mocks"
	"testing"
)

func TestUserUsecase_SignIn_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockUserRepositoryInterface(mockCtrl)

	UserUsecase := UserUsecase{
		Clients:         nil,
		Db:              dbMock,
		LoggerInterface: &models.Logger{Logger: &logrus.Entry{}},
		Sanitizer:       bluemonday.NewPolicy(),
	}

	pass, _ := UserUsecase.HashPassword("12345678")

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

	returnUser, code, err := UserUsecase.SignIn(user)

	assert.Equal(t, 200, code)
	assert.Equal(t, nil, err)
	assert.Equal(t, user1, returnUser)
}

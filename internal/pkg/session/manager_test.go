package session

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"server/internal/pkg/models"
	"server/internal/pkg/session/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSetSession(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockSessionRepositoryInterface(mockCtrl)

	SessionsManagerTest := SessionsManager{
		DB:     dbMock,
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	rw := httptest.NewRecorder()

	userid := 1

	dbMock.EXPECT().AddCookie(userid, gomock.Any()).Return(nil)

	err := SessionsManagerTest.SetSession(rw, userid)

	assert.Equal(t, nil, err)
}

func TestAddCookieError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockSessionRepositoryInterface(mockCtrl)

	SessionsManagerTest := SessionsManager{
		DB:     dbMock,
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	rw := httptest.NewRecorder()

	userid := 1

	dbMock.EXPECT().AddCookie(userid, gomock.Any()).Return(errors.New("Some error"))

	err := SessionsManagerTest.SetSession(rw, userid)

	assert.NotEqual(t, nil, err)
}

func TestDeleteSession(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockSessionRepositoryInterface(mockCtrl)

	SessionsManagerTest := SessionsManager{
		DB:     dbMock,
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	cookie := http.Cookie{
		Name:  "token",
		Value: "123",
	}

	dbMock.EXPECT().DeleteCookie("123").Return(nil)

	err := SessionsManagerTest.DeleteSession(&cookie)

	assert.Equal(t, nil, err)
}

func TestDeleteSessionDeleteCookieError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockSessionRepositoryInterface(mockCtrl)

	SessionsManagerTest := SessionsManager{
		DB:     dbMock,
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	cookie := http.Cookie{
		Name:  "token",
		Value: "123",
	}

	dbMock.EXPECT().DeleteCookie("123").Return(errors.New("Some error"))

	err := SessionsManagerTest.DeleteSession(&cookie)

	assert.NotEqual(t, nil, err)
}

func TestGetIdFromContext(t *testing.T) {
	SessionsManagerTest := SessionsManager{
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	req := &http.Request{}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		1,
	)

	id, ok := SessionsManagerTest.GetIdFromContext(ctx)

	assert.Equal(t, 1, id)
	assert.Equal(t, true, ok)
}

func TestGetIdFromContextError(t *testing.T) {
	SessionsManagerTest := SessionsManager{
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	req := &http.Request{}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		"some key",
		1,
	)

	_, ok := SessionsManagerTest.GetIdFromContext(ctx)

	assert.Equal(t, false, ok)
}

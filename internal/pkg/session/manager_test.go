package session

import (
	"context"
	"errors"
	"net/http"
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

	userid := 1

	dbMock.EXPECT().AddCookie(userid, gomock.Any()).Return(nil)

	_, err := SessionsManagerTest.SetSession(userid)

	assert.Equal(t, nil, err)
}

func TestAddCookieError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockSessionRepositoryInterface(mockCtrl)

	SessionsManagerTest := SessionsManager{
		DB:     dbMock,
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	userid := 1

	dbMock.EXPECT().AddCookie(userid, gomock.Any()).Return(errors.New("Some error"))

	_, err := SessionsManagerTest.SetSession(userid)

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

func TestCheckSession(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockSessionRepositoryInterface(mockCtrl)

	SessionsManagerTest := SessionsManager{
		DB:     dbMock,
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	userid := 1
	tokens := []string{"123", "123"}

	dbMock.EXPECT().GetCookie(gomock.Any()).Return(userid, nil)

	id, ok := SessionsManagerTest.CheckSession(tokens)

	assert.Equal(t, true, ok)
	assert.Equal(t, id, userid)
}

func TestCheckSession_GetCookie_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockSessionRepositoryInterface(mockCtrl)

	SessionsManagerTest := SessionsManager{
		DB:     dbMock,
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	userid := 1
	tokens := []string{"123", "123"}

	dbMock.EXPECT().GetCookie(gomock.Any()).Return(userid, errors.New("Some error"))

	id, ok := SessionsManagerTest.CheckSession(tokens)

	assert.Equal(t, false, ok)
	assert.Equal(t, -1, id)
}

func TestCheckSession_NotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockSessionRepositoryInterface(mockCtrl)

	SessionsManagerTest := SessionsManager{
		DB:     dbMock,
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	tokens := []string{"123"}

	dbMock.EXPECT().GetCookie(gomock.Any()).Return(-1, nil)

	id, ok := SessionsManagerTest.CheckSession(tokens)

	assert.Equal(t, false, ok)
	assert.Equal(t, id, -1)
}

func TestDeleteSessionById(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockSessionRepositoryInterface(mockCtrl)

	SessionsManagerTest := SessionsManager{
		DB:     dbMock,
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	dbMock.EXPECT().DeleteCookie(gomock.Any(), gomock.Any()).Return(nil)

	err := SessionsManagerTest.DeleteSessionById(1)

	assert.Equal(t, err, nil)
}

func TestDeleteSessionById_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockSessionRepositoryInterface(mockCtrl)

	SessionsManagerTest := SessionsManager{
		DB:     dbMock,
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	dbMock.EXPECT().DeleteCookie(gomock.Any(), gomock.Any()).Return(errors.New("Some error"))

	err := SessionsManagerTest.DeleteSessionById(1)

	assert.NotEqual(t, err, nil)
}

func TestDeleteSessionByToken(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockSessionRepositoryInterface(mockCtrl)

	SessionsManagerTest := SessionsManager{
		DB:     dbMock,
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	dbMock.EXPECT().DeleteCookie(gomock.Any(), gomock.Any()).Return(nil)

	err := SessionsManagerTest.DeleteSessionByToken("1")

	assert.Equal(t, err, nil)
}

func TestDeleteSessionByToken_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	dbMock := mocks.NewMockSessionRepositoryInterface(mockCtrl)

	SessionsManagerTest := SessionsManager{
		DB:     dbMock,
		Logger: &models.Logger{Logger: logrus.New().WithField("test", "test")},
	}

	dbMock.EXPECT().DeleteCookie(gomock.Any(), gomock.Any()).Return(errors.New("Some error"))

	err := SessionsManagerTest.DeleteSessionByToken("1")

	assert.NotEqual(t, err, nil)
}

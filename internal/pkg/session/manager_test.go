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

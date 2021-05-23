package delivery

import (
	"context"
	"errors"
	"net/http"
	session_proto "server/internal/auth_server/delivery/session"
	"server/internal/pkg/models"
	"server/internal/pkg/session/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	useCaseMock := mocks.NewMockSessionManagerInterface(mockCtrl)

	AuthServer := AuthServer{
		Usecase: useCaseMock,
	}

	userid := 1

	req := &http.Request{
		Method: "POST",
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userid,
	)

	useCaseMock.EXPECT().SetSession(gomock.Any()).Return("token", nil)

	_, err := AuthServer.Create(ctx, &session_proto.SessionId{Id: int32(userid)})

	assert.Equal(t, nil, err)
}

func TestCreate_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	useCaseMock := mocks.NewMockSessionManagerInterface(mockCtrl)

	AuthServer := AuthServer{
		Usecase: useCaseMock,
	}

	userid := 1

	req := &http.Request{
		Method: "POST",
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userid,
	)

	useCaseMock.EXPECT().SetSession(gomock.Any()).Return("token", errors.New("Some error"))

	_, err := AuthServer.Create(ctx, &session_proto.SessionId{Id: int32(userid)})

	assert.NotEqual(t, nil, err)
}

func TestCheck(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	useCaseMock := mocks.NewMockSessionManagerInterface(mockCtrl)

	AuthServer := AuthServer{
		Usecase: useCaseMock,
	}

	userid := 1

	req := &http.Request{
		Method: "POST",
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userid,
	)

	useCaseMock.EXPECT().CheckSession(gomock.Any()).Return(userid, true)

	_, err := AuthServer.Check(ctx, &session_proto.SessionToken{Token: "token"})

	assert.Equal(t, nil, err)
}

func TestCheck_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	useCaseMock := mocks.NewMockSessionManagerInterface(mockCtrl)

	AuthServer := AuthServer{
		Usecase: useCaseMock,
	}

	userid := 1

	req := &http.Request{
		Method: "POST",
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userid,
	)

	useCaseMock.EXPECT().CheckSession(gomock.Any()).Return(userid, false)

	_, err := AuthServer.Check(ctx, &session_proto.SessionToken{Token: "token"})

	assert.NotEqual(t, nil, err)
}

func TestDelete(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	useCaseMock := mocks.NewMockSessionManagerInterface(mockCtrl)

	AuthServer := AuthServer{
		Usecase: useCaseMock,
	}

	userid := 1

	req := &http.Request{
		Method: "POST",
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx,
		models.CtxUserId,
		userid,
	)

	useCaseMock.EXPECT().DeleteSessionById(gomock.Any()).Return(nil)

	_, err := AuthServer.Delete(ctx, &session_proto.SessionId{Id: int32(userid)})

	assert.Equal(t, nil, err)
}

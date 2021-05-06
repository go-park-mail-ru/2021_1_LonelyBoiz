package delivery

import (
	"net/http"
	sessionMocks "server/internal/auth_server/delivery/session/mocks"
	"server/internal/pkg/models"
	mock_usecase "server/internal/pkg/user/usecase/mocks"
	"testing"

	"server/internal/pkg/user/usecase"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	userUseCase := usecase.UserUsecase{}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().ProtoUser2User(userUseCase.User2ProtoUser(user)).Return(user)
	userUseCaseMock.EXPECT().CreateNewUser(user).Return(user, 200, nil)
	userUseCaseMock.EXPECT().User2ProtoUser(user).Return(userUseCase.User2ProtoUser(user))
	sessionManagerMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, nil)

	_, err := server.CreateUser(ctx, userUseCase.User2ProtoUser(user))

	assert.Equal(t, err, nil)
}

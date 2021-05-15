package delivery

import (
	"errors"
	"net/http"
	session_proto "server/internal/auth_server/delivery/session"
	sessionMocks "server/internal/auth_server/delivery/session/mocks"
	"server/internal/pkg/models"
	mock_usecase "server/internal/pkg/user/usecase/mocks"
	user_proto "server/internal/user_server/delivery/proto"
	"testing"

	"server/internal/pkg/user/usecase"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateUser(t *testing.T) {
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

func TestCreateUser_DB_Error(t *testing.T) {
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
	userUseCaseMock.EXPECT().CreateNewUser(user).Return(user, 500, errors.New("Some error"))

	_, err := server.CreateUser(ctx, userUseCase.User2ProtoUser(user))

	assert.NotEqual(t, err, nil)
}

func TestCreateUser_Session_Error(t *testing.T) {
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

	newErr := status.Error(codes.Code(500), "Some error")

	userUseCaseMock.EXPECT().ProtoUser2User(userUseCase.User2ProtoUser(user)).Return(user)
	userUseCaseMock.EXPECT().CreateNewUser(user).Return(user, 200, nil)
	sessionManagerMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("Some error"))

	_, err := server.CreateUser(ctx, userUseCase.User2ProtoUser(user))

	assert.Equal(t, err, newErr)
}

func TestDeleteUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().CheckIds(ctx).Return(user.Id, 200, nil)
	userUseCaseMock.EXPECT().DeleteUser(user.Id).Return(nil)
	sessionManagerMock.EXPECT().Delete(ctx, gomock.Any()).Return(&session_proto.Nothing{}, nil)

	_, err := server.DeleteUser(ctx, &user_proto.UserNothing{})

	assert.Equal(t, err, nil)
}

func TestDeleteUser_CheckIds_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().CheckIds(ctx).Return(-1, 500, errors.New("Some error"))

	_, err := server.DeleteUser(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestDeleteUser_DB_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().CheckIds(ctx).Return(user.Id, 200, nil)
	userUseCaseMock.EXPECT().DeleteUser(user.Id).Return(errors.New("Some error"))

	_, err := server.DeleteUser(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestDeleteUser_Sessions_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().CheckIds(ctx).Return(user.Id, 200, nil)
	userUseCaseMock.EXPECT().DeleteUser(user.Id).Return(nil)
	sessionManagerMock.EXPECT().Delete(ctx, gomock.Any()).Return(nil, errors.New("Some error"))

	_, err := server.DeleteUser(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestChangeUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}
	protoUser := user_proto.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().CheckIds(ctx).Return(user.Id, 200, nil)
	userUseCaseMock.EXPECT().ProtoUser2User(&protoUser).Return(user)
	userUseCaseMock.EXPECT().ChangeUserInfo(user, user.Id).Return(user, 200, nil)
	userUseCaseMock.EXPECT().User2ProtoUser(user).Return(&protoUser)

	_, err := server.ChangeUser(ctx, &protoUser)

	assert.Equal(t, err, nil)
}

func TestChangeUser_CheckIds_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}
	protoUser := user_proto.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().CheckIds(ctx).Return(user.Id, 200, errors.New("Some error"))

	_, err := server.ChangeUser(ctx, &protoUser)

	assert.NotEqual(t, err, nil)
}

func TestChangeUser_ChangeUserInfo_ValidationError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}
	protoUser := user_proto.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().CheckIds(ctx).Return(user.Id, 200, nil)
	userUseCaseMock.EXPECT().ProtoUser2User(&protoUser).Return(user)
	userUseCaseMock.EXPECT().ChangeUserInfo(user, user.Id).Return(user, 400, errors.New("Some error"))

	_, err := server.ChangeUser(ctx, &protoUser)

	assert.NotEqual(t, err, nil)
}

func TestChangeUser_ChangeUserInfo_InternalError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}
	protoUser := user_proto.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().CheckIds(ctx).Return(user.Id, 200, nil)
	userUseCaseMock.EXPECT().ProtoUser2User(&protoUser).Return(user)
	userUseCaseMock.EXPECT().ChangeUserInfo(user, user.Id).Return(user, 500, nil)

	_, err := server.ChangeUser(ctx, &protoUser)

	assert.NotEqual(t, err, nil)
}

func TestCheckUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}
	protoUser := user_proto.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	loginUser := user_proto.UserLogin{
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().SignIn(gomock.Any()).Return(user, 200, nil)
	sessionManagerMock.EXPECT().Create(ctx, &session_proto.SessionId{Id: int32(user.Id)}).Return(&session_proto.SessionToken{}, nil)
	userUseCaseMock.EXPECT().User2ProtoUser(user).Return(&protoUser)

	_, err := server.CheckUser(ctx, &loginUser)

	assert.Equal(t, err, nil)
}

func TestCheckUser_SignIn_InternalError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	loginUser := user_proto.UserLogin{
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().SignIn(gomock.Any()).Return(user, 500, errors.New("Some error"))

	_, err := server.CheckUser(ctx, &loginUser)

	assert.NotEqual(t, err, nil)
}

func TestCheckUser_SignIn_ValidationError(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	loginUser := user_proto.UserLogin{
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().SignIn(gomock.Any()).Return(user, 400, errors.New("Some error"))

	_, err := server.CheckUser(ctx, &loginUser)

	assert.NotEqual(t, err, nil)
}

func TestCheckUser_Sessions_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	loginUser := user_proto.UserLogin{
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().SignIn(gomock.Any()).Return(user, 200, nil)
	sessionManagerMock.EXPECT().Create(ctx, &session_proto.SessionId{Id: int32(user.Id)}).Return(nil, errors.New("Some error"))

	_, err := server.CheckUser(ctx, &loginUser)

	assert.NotEqual(t, err, nil)
}

func TestGetUserById(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}
	protoUser := user_proto.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetUserInfoById(user.Id).Return(user, nil)
	userUseCaseMock.EXPECT().User2ProtoUser(user).Return(&protoUser)

	_, err := server.GetUserById(ctx, &user_proto.UserNothing{})

	assert.Equal(t, err, nil)
}

func TestGetUserById_GetUserInfoById_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlId").Return(user.Id, false)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetUserInfoById(user.Id).Return(user, errors.New("Some error"))

	_, err := server.GetUserById(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestGetUserById_GetParamDromContext_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlId").Return(user.Id, false)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, false)

	_, err := server.GetUserById(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

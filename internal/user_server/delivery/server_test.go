package delivery

import (
	"errors"
	"net/http"
	session_proto "server/internal/auth_server/delivery/session"
	sessionMocks "server/internal/auth_server/delivery/session/mocks"
	chat_usecase "server/internal/pkg/chat/usecase/mocks"
	message_usecase "server/internal/pkg/message/usecase/mocks"
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

func TestBlockSecretAlbum(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(gomock.Any(), gomock.Any()).Return(1, true)
	userUseCaseMock.EXPECT().GetParamFromContext(gomock.Any(), gomock.Any()).Return(1, true)
	userUseCaseMock.EXPECT().BlockSecreteAlbum(1, 1).Return(200, nil)

	_, err := server.BlockSecretAlbum(ctx, &user_proto.UserNothing{})

	assert.Equal(t, err, nil)
}

func TestBlockSecretAlbum_cookieId_NotOk(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(gomock.Any(), gomock.Any()).Return(1, false)

	_, err := server.BlockSecretAlbum(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestBlockSecretAlbum_getterId_NotOk(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(gomock.Any(), gomock.Any()).Return(1, true)
	userUseCaseMock.EXPECT().GetParamFromContext(gomock.Any(), gomock.Any()).Return(1, false)

	_, err := server.BlockSecretAlbum(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestBlockSecretAlbum_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(gomock.Any(), gomock.Any()).Return(1, true)
	userUseCaseMock.EXPECT().GetParamFromContext(gomock.Any(), gomock.Any()).Return(1, true)
	userUseCaseMock.EXPECT().BlockSecreteAlbum(1, 1).Return(500, errors.New("Some error"))

	_, err := server.BlockSecretAlbum(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestDeleteChat(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(gomock.Any(), gomock.Any()).Return(1, true)
	userUseCaseMock.EXPECT().DeleteChat(1).Return(nil)

	_, err := server.DeleteChat(ctx, &user_proto.UserNothing{})

	assert.Equal(t, err, nil)
}

func TestDeleteChat_false(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(gomock.Any(), gomock.Any()).Return(-1, false)

	_, err := server.DeleteChat(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestDeleteChat_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(gomock.Any(), gomock.Any()).Return(1, true)
	userUseCaseMock.EXPECT().DeleteChat(1).Return(errors.New("Some error"))

	_, err := server.DeleteChat(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

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

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().ProtoUser2User(userUseCase.User2ProtoUser(user)).Return(user)
	userUseCaseMock.EXPECT().CreateNewUser(user).Return(user, 200, nil)
	userUseCaseMock.EXPECT().User2ProtoUser(user).Return(userUseCase.User2ProtoUser(user))
	sessionManagerMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, nil)
	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()

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

	userUseCaseMock.EXPECT().LogInfo(gomock.Any()).Return()
	userUseCaseMock.EXPECT().LogError(gomock.Any()).Return()
	userUseCaseMock.EXPECT().ProtoUser2User(userUseCase.User2ProtoUser(user)).Return(user)
	userUseCaseMock.EXPECT().CreateNewUser(user).Return(user, 200, nil)
	sessionManagerMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, newErr)

	_, err := server.CreateUser(ctx, userUseCase.User2ProtoUser(user))

	assert.NotEqual(t, err, nil)
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

func TestCreateFeed(t *testing.T) {
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

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlCount").Return(20, true)
	userUseCaseMock.EXPECT().CreateFeed(user.Id, 20).Return([]int{1, 2}, 200, nil)

	_, err := server.CreateFeed(ctx, &user_proto.UserNothing{})

	assert.Equal(t, err, nil)
}

func TestCreateFeed_GetCookie_Error(t *testing.T) {
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

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, false)

	_, err := server.CreateFeed(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestCreateFeed_GetCount_Error(t *testing.T) {
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

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlCount").Return(20, false)

	_, err := server.CreateFeed(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestCreateFeed_InternalError(t *testing.T) {
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

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlCount").Return(20, true)
	userUseCaseMock.EXPECT().CreateFeed(user.Id, 20).Return([]int{1, 2}, 500, errors.New("Some error"))

	_, err := server.CreateFeed(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestCreateFeed_ValidationError(t *testing.T) {
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

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlCount").Return(20, true)
	userUseCaseMock.EXPECT().CreateFeed(user.Id, 20).Return([]int{1, 2}, 400, errors.New("Some error"))

	_, err := server.CreateFeed(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestCreateChat(t *testing.T) {
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

	chat := models.Chat{
		ChatId:      1,
		PartnerId:   2,
		PartnerName: "Name",
		Photos:      []string{},
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().CreateChat(user.Id, models.Like{}).Return(chat, 200, nil)
	userUseCaseMock.EXPECT().Photos2ProtoPhotos(gomock.Any()).Return(chat.Photos)

	_, err := server.CreateChat(ctx, &user_proto.Like{})

	assert.Equal(t, err, nil)
}

func TestCreateChat_GetCookie_Error(t *testing.T) {
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

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, false)

	_, err := server.CreateChat(ctx, &user_proto.Like{})

	assert.NotEqual(t, err, nil)
}

func TestCreateChat_204(t *testing.T) {
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

	chat := models.Chat{
		ChatId:      1,
		PartnerId:   2,
		PartnerName: "Name",
		Photos:      []string{},
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().CreateChat(user.Id, models.Like{}).Return(chat, 204, nil)

	_, err := server.CreateChat(ctx, &user_proto.Like{})

	assert.NotEqual(t, err, nil)
}

func TestCreateChat_Error(t *testing.T) {
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

	chat := models.Chat{
		ChatId:      1,
		PartnerId:   2,
		PartnerName: "Name",
		Photos:      []string{},
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().CreateChat(user.Id, models.Like{}).Return(chat, 500, errors.New("Some error"))

	_, err := server.CreateChat(ctx, &user_proto.Like{})

	assert.NotEqual(t, err, nil)
}

func TestCreateChat_Validation_Error(t *testing.T) {
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

	chat := models.Chat{
		ChatId:      1,
		PartnerId:   2,
		PartnerName: "Name",
		Photos:      []string{},
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().CreateChat(user.Id, models.Like{}).Return(chat, 400, errors.New("Some error"))

	_, err := server.CreateChat(ctx, &user_proto.Like{})

	assert.NotEqual(t, err, nil)
}

func TestGetChats(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
		ChatUsecase: chatUseCaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	chat := models.Chat{
		ChatId:      1,
		PartnerId:   2,
		PartnerName: "Name",
		Photos:      []string{},
	}

	chats := []models.Chat{chat}

	limit := 20
	offset := 20

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlCount").Return(limit, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlOffset").Return(offset, true)

	chatUseCaseMock.EXPECT().GetChat(user.Id, offset, limit).Return(chats, nil)
	chatUseCaseMock.EXPECT().Chat2ProtoChat(chat).Return(&user_proto.Chat{})

	_, err := server.GetChats(ctx, &user_proto.UserNothing{})

	assert.Equal(t, err, nil)
}

func TestGetChats_GetCookie_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
		ChatUsecase: chatUseCaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, false)

	_, err := server.GetChats(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestGetChats_GetCount_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
		ChatUsecase: chatUseCaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	limit := 20

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlCount").Return(limit, false)

	_, err := server.GetChats(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestGetChats_GetOffset_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
		ChatUsecase: chatUseCaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	limit := 20
	offset := 20

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlCount").Return(limit, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlOffset").Return(offset, false)

	_, err := server.GetChats(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestGetChats_GetChat_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase: userUseCaseMock,
		Sessions:    sessionManagerMock,
		ChatUsecase: chatUseCaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	chat := models.Chat{
		ChatId:      1,
		PartnerId:   2,
		PartnerName: "Name",
		Photos:      []string{},
	}

	chats := []models.Chat{chat}

	limit := 20
	offset := 20

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlCount").Return(limit, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlOffset").Return(offset, true)

	chatUseCaseMock.EXPECT().GetChat(user.Id, offset, limit).Return(chats, errors.New("Some error"))

	_, err := server.GetChats(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestGetMessages(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	message := models.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	messages := []models.Message{message}

	limit := 20
	offset := 20

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlCount").Return(limit, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlOffset").Return(offset, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlChatId").Return(message.ChatId, true)

	messageUsecaseMock.EXPECT().ManageMessage(user.Id, message.ChatId, limit, offset).Return(messages, 200, nil)
	messageUsecaseMock.EXPECT().Message2ProtoMessage(message).Return(&user_proto.Message{})

	_, err := server.GetMessages(ctx, &user_proto.UserNothing{})

	assert.Equal(t, err, nil)
}

func TestGetMessages_getCookie_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, false)

	_, err := server.GetMessages(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestGetMessages_GetCount_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	limit := 20

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlCount").Return(limit, false)

	_, err := server.GetMessages(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestGetMessages_GetOffset_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	limit := 20
	offset := 20

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlCount").Return(limit, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlOffset").Return(offset, false)

	_, err := server.GetMessages(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestGetMessages_GetChatId_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	message := models.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	limit := 20
	offset := 20

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlCount").Return(limit, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlOffset").Return(offset, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlChatId").Return(message.ChatId, false)

	_, err := server.GetMessages(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestGetMessages_ManageMessage_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	message := models.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	messages := []models.Message{message}

	limit := 20
	offset := 20

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlCount").Return(limit, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlOffset").Return(offset, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlChatId").Return(message.ChatId, true)

	messageUsecaseMock.EXPECT().ManageMessage(user.Id, message.ChatId, limit, offset).Return(messages, 500, errors.New("Some error"))

	_, err := server.GetMessages(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestCreateMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	message := models.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	protoMessage := user_proto.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlChatId").Return(message.ChatId, true)

	messageUsecaseMock.EXPECT().ProtoMessage2Message(gomock.Any()).Return(message)
	messageUsecaseMock.EXPECT().Message2ProtoMessage(gomock.Any()).Return(&protoMessage)

	messageUsecaseMock.EXPECT().CreateMessage(message, message.ChatId, user.Id).Return(message, 200, nil)

	_, err := server.CreateMessage(ctx, &protoMessage)

	assert.Equal(t, err, nil)
}

func TestCreateMessage_GetCookie_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	protoMessage := user_proto.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, false)

	_, err := server.CreateMessage(ctx, &protoMessage)

	assert.NotEqual(t, err, nil)
}

func TestCreateMessage_GetChat_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	message := models.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	protoMessage := user_proto.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlChatId").Return(message.ChatId, false)

	_, err := server.CreateMessage(ctx, &protoMessage)

	assert.NotEqual(t, err, nil)
}

func TestCreateMessage_CreateMessage_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	message := models.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	protoMessage := user_proto.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlChatId").Return(message.ChatId, true)

	messageUsecaseMock.EXPECT().ProtoMessage2Message(gomock.Any()).Return(message)

	messageUsecaseMock.EXPECT().CreateMessage(message, message.ChatId, user.Id).Return(message, 400, errors.New("Some error"))

	_, err := server.CreateMessage(ctx, &protoMessage)

	assert.NotEqual(t, err, nil)
}

func TestChangeMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	message := models.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	protoMessage := user_proto.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlMessageId").Return(message.ChatId, true)

	messageUsecaseMock.EXPECT().ProtoMessage2Message(&protoMessage).Return(message)
	messageUsecaseMock.EXPECT().ChangeMessage(user.Id, message.MessageId, message).Return(message, 204, nil)
	messageUsecaseMock.EXPECT().Message2ProtoMessage(gomock.Any()).Return(&protoMessage)

	_, err := server.ChangeMessage(ctx, &protoMessage)

	assert.Equal(t, err, nil)
}

func TestChangeMessage_GetCookie_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	protoMessage := user_proto.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, false)

	_, err := server.ChangeMessage(ctx, &protoMessage)

	assert.NotEqual(t, err, nil)
}

func TestChangeMessage_GetMessageId_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	message := models.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	protoMessage := user_proto.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlMessageId").Return(message.ChatId, false)

	_, err := server.ChangeMessage(ctx, &protoMessage)

	assert.NotEqual(t, err, nil)
}

func TestChangeMessage_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	message := models.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	protoMessage := user_proto.Message{
		MessageId: 1,
		AuthorId:  1,
		ChatId:    1,
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "urlMessageId").Return(message.ChatId, true)

	messageUsecaseMock.EXPECT().ProtoMessage2Message(&protoMessage).Return(message)
	messageUsecaseMock.EXPECT().ChangeMessage(user.Id, message.MessageId, message).Return(message, 500, errors.New("some error"))

	_, err := server.ChangeMessage(ctx, &protoMessage)

	assert.NotEqual(t, err, nil)
}

func TestAddToSecretAlbum(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
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

	userUseCaseMock.EXPECT().ProtoPhotos2Photos(gomock.Any()).Return([]string{})
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().AddToSecreteAlbum(user.Id, gomock.Any()).Return(200, nil)

	_, err := server.AddToSecreteAlbum(ctx, &protoUser)

	assert.Equal(t, err, nil)
}

func TestAddToSecretAlbum_GetCookie_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
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

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, false)

	_, err := server.AddToSecreteAlbum(ctx, &protoUser)

	assert.NotEqual(t, err, nil)
}

func TestAddToSecretAlbum_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
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

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().ProtoPhotos2Photos(gomock.Any()).Return([]string{})
	userUseCaseMock.EXPECT().AddToSecreteAlbum(user.Id, gomock.Any()).Return(500, errors.New("Some error"))

	_, err := server.AddToSecreteAlbum(ctx, &protoUser)

	assert.NotEqual(t, err, nil)
}

func TestUnlockSecretAlbum(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "getterId").Return(user.Id, true)
	userUseCaseMock.EXPECT().UnblockSecreteAlbum(user.Id, user.Id).Return(200, nil)

	_, err := server.UnlockSecretAlbum(ctx, &user_proto.UserNothing{})

	assert.Equal(t, err, nil)
}

func TestUnlockSecretAlbum_GetCookie_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, false)

	_, err := server.UnlockSecretAlbum(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestUnlockSecretAlbum_GetId_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "getterId").Return(user.Id, false)

	_, err := server.UnlockSecretAlbum(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestUnlockSecretAlbum_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "getterId").Return(user.Id, true)
	userUseCaseMock.EXPECT().UnblockSecreteAlbum(user.Id, user.Id).Return(200, errors.New("Some error"))

	_, err := server.UnlockSecretAlbum(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestGetSecreteAlbum(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
		Photos:         []string{},
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "ownerId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetSecreteAlbum(user.Id, user.Id).Return([]string{}, 200, nil)
	userUseCaseMock.EXPECT().Photos2ProtoPhotos(gomock.Any()).Return(user.Photos)

	_, err := server.GetSecreteAlbum(ctx, &user_proto.UserNothing{})

	assert.Equal(t, err, nil)
}

func TestGetSecreteAlbum_GetOwnerId_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
		Photos:         []string{},
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "ownerId").Return(user.Id, false)

	_, err := server.GetSecreteAlbum(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestGetSecreteAlbum_GetCookie_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
		Photos:         []string{},
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "ownerId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, false)

	_, err := server.GetSecreteAlbum(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

func TestGetSecreteAlbum_Error(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	userUseCaseMock := mock_usecase.NewMockUserUseCaseInterface(mockCtrl)
	sessionManagerMock := sessionMocks.NewMockAuthCheckerClient(mockCtrl)
	chatUseCaseMock := chat_usecase.NewMockChatUsecaseInterface(mockCtrl)
	messageUsecaseMock := message_usecase.NewMockMessageUsecaseInterface(mockCtrl)

	server := UserServer{
		UserUsecase:    userUseCaseMock,
		Sessions:       sessionManagerMock,
		ChatUsecase:    chatUseCaseMock,
		MessageUsecase: messageUsecaseMock,
	}

	user := models.User{
		Id:             1,
		Email:          "windes",
		Password:       "12345678",
		SecondPassword: "12345678",
		Photos:         []string{},
	}

	req := &http.Request{}

	ctx := req.Context()

	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "ownerId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetParamFromContext(ctx, "cookieId").Return(user.Id, true)
	userUseCaseMock.EXPECT().GetSecreteAlbum(user.Id, user.Id).Return([]string{}, 500, errors.New("Some error"))

	_, err := server.GetSecreteAlbum(ctx, &user_proto.UserNothing{})

	assert.NotEqual(t, err, nil)
}

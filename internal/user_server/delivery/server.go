package delivery

import (
	"golang.org/x/net/context"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	session_proto2 "server/internal/auth_server/delivery/session"
	model "server/internal/pkg/models"
	"server/internal/pkg/user/usecase"
	userProto "server/internal/user_server/delivery/proto"
)

type UserServer struct {
	userProto.UnimplementedUserServiceServer
	Usecase  usecase.UserUseCaseInterface
	Sessions session_proto2.AuthCheckerClient
}

func (u UserServer) CreateUser(ctx context.Context, user *userProto.User) (*userProto.UserResponse, error) {
	nUser, ok := u.Usecase.ProtoUser2User(user)
	if !ok {
		return &userProto.UserResponse{}, status.Error(500, "")
	}

	newUser, code, responseError := u.Usecase.CreateNewUser(nUser)
	if code != 200 {
		return &userProto.UserResponse{}, status.Error(codes.Code(code), responseError.Error())
	}

	token, err := u.Sessions.Create(ctx, &session_proto2.SessionId{Id: int32(newUser.Id)})
	if err != nil {
		return &userProto.UserResponse{}, status.Error(codes.Code(code), err.Error())
	}

	protoUser, ok := u.Usecase.User2ProtoUser(newUser)
	if !ok {
		return &userProto.UserResponse{}, status.Error(500, "")
	}

	return &userProto.UserResponse{
		User:  protoUser,
		Token: token.GetToken(),
	}, nil
}

func (u UserServer) DeleteUser(ctx context.Context, nothing *userProto.UserNothing) (*userProto.UserNothing, error) {
	id, code, err := u.Usecase.CheckIds(ctx)
	if err != nil {
		return &userProto.UserNothing{}, status.Error(codes.Code(code), err.Error())
	}

	err = u.Usecase.DeleteUser(id)
	if err != nil {
		return &userProto.UserNothing{}, status.Error(500, "")
	}

	_, err = u.Sessions.Delete(ctx, &session_proto2.SessionId{Id: int32(id)})
	if err != nil {
		return &userProto.UserNothing{}, status.Error(500, "")
	}

	return &userProto.UserNothing{}, nil
}

func (u UserServer) ChangeUser(ctx context.Context, user *userProto.User) (*userProto.User, error) {
	id, code, err := u.Usecase.CheckIds(ctx)
	if err != nil {
		return &userProto.User{}, status.Error(codes.Code(code), err.Error())
	}

	nUser, ok := u.Usecase.ProtoUser2User(user)
	if !ok {
		return &userProto.User{}, status.Error(500, "")
	}

	newUser, code, err := u.Usecase.ChangeUserInfo(nUser, id)
	if code != 200 {
		if err != nil {
			return &userProto.User{}, status.Error(codes.Code(code), err.Error())
		}
		return &userProto.User{}, status.Error(codes.Code(code), "")
	}

	protoUser, ok := u.Usecase.User2ProtoUser(newUser)
	if !ok {
		return &userProto.User{}, status.Error(500, "")
	}

	if protoUser == nil {
		return &userProto.User{}, status.Error(500, "")
	}

	return protoUser, nil
}

func (UserServer) CheckUser(ctx context.Context, login *userProto.UserLogin) (*userProto.User, error) {
	panic("implement me")
}

func (u UserServer) GetUserById(ctx context.Context, nothing *userProto.UserNothing) (*userProto.User, error) {
	id, ok := u.Usecase.GetParamFromContext(ctx, "urlId")
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		return &userProto.User{}, status.Error(403, response.Error())
	}

	userInfo, err := u.Usecase.GetUserInfoById(id)
	if err != nil {
		response := model.ErrorResponse{Err: "Пользователь не найден"}
		return &userProto.User{}, status.Error(401, response.Error())
	}

	protoUser, ok := u.Usecase.User2ProtoUser(userInfo)
	if !ok {
		return &userProto.User{}, status.Error(500, "")
	}

	return protoUser, nil
}

func (UserServer) CreateFeed(ctx context.Context, nothing *userProto.UserNothing) (*userProto.Feed, error) {
	panic("implement me")
}

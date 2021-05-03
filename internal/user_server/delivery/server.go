package delivery

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"log"
	session_proto2 "server/internal/auth_server/delivery/session"
	"server/internal/pkg/user/usecase"
	userProto "server/internal/user_server/delivery/proto"
	"strconv"
)

type UserServer struct {
	userProto.UnimplementedUserServiceServer
	Usecase  usecase.UserUseCaseInterface
	Sessions session_proto2.AuthCheckerClient
}

func (u UserServer) CreateUser(ctx context.Context, user *userProto.User) (*userProto.UserResponse, error) {
	newUser, code, responseError := u.Usecase.CreateNewUser(u.Usecase.ProtoUser2User(user))
	if code != 200 {
		return &userProto.UserResponse{
			Code:  int32(code),
			Token: "",
			User:  nil,
		}, responseError
	}

	//token, err := u.Sessions.Create(ctx, &session_proto2.SessionId{Id: int32(newUser.Id)})
	//if err != nil {
	//	return &userProto.UserResponse{
	//		Code: 500,
	//		User: nil,
	//	}, err
	//}

	return &userProto.UserResponse{
		Code:  200,
		User:  u.Usecase.User2ProtoUser(newUser),
		Token: "",
	}, nil
}

func (UserServer) DeleteUser(ctx context.Context, nothing *userProto.UserNothing) (*userProto.UserNothing, error) {
	panic("implement me")
}

func (UserServer) ChangeUser(ctx context.Context, user *userProto.User) (*userProto.UserResponse, error) {
	panic("implement me")
}

func (UserServer) CheckUser(ctx context.Context, login *userProto.UserLogin) (*userProto.User, error) {
	panic("implement me")
}

func (u UserServer) GetUserById(ctx context.Context, nothing *userProto.UserNothing) (*userProto.User, error) {
	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		panic("context")
	}

	ids := data.Get("id")

	log.Println(ids)

	if len(ids) == 0 {
		panic("len")
	}
	id, err := strconv.Atoi(ids[0])
	if err != nil {
		panic(err)
	}

	userInfo, err := u.Usecase.GetUserInfoById(id)
	if err != nil {
		panic(err)
	}

	return &userProto.User{
		Id:             int32(id),
		Email:          userInfo.Email,
		Name:           userInfo.Name,
		Birthday:       userInfo.Birthday,
		Description:    userInfo.Description,
		City:           userInfo.City,
		Instagram:      userInfo.Instagram,
		Sex:            userInfo.Sex,
		DatePreference: userInfo.DatePreference,
		IsDeleted:      userInfo.IsDeleted,
		IsActive:       userInfo.IsActive,
		CaptchaToken:   userInfo.CaptchaToken,
	}, nil
}

func (UserServer) CreateFeed(ctx context.Context, nothing *userProto.UserNothing) (*userProto.Feed, error) {
	panic("implement me")
}

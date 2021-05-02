package delivery

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"log"
	"server/internal/pkg/user/usecase"
	userProto "server/internal/user_server/delivery/proto"
	"strconv"
)

type UserServer struct {
	userProto.UnimplementedUserServiceServer
	Usecase usecase.UserUseCaseInterface
}

func (UserServer) CreateUser(ctx context.Context, user *userProto.User) (*userProto.User, error) {
	panic("implement me")
}

func (UserServer) DeleteUser(ctx context.Context, nothing *userProto.UserNothing) (*userProto.UserNothing, error) {
	panic("implement me")
}

func (UserServer) ChangeUser(ctx context.Context, user *userProto.User) (*userProto.User, error) {
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

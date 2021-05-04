package delivery

import (
	"golang.org/x/net/context"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	session_proto2 "server/internal/auth_server/delivery/session"
	model "server/internal/pkg/models"
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
			User: nil,
		}, status.Error(codes.Code(code), responseError.Error())
	}

	//token, err := u.Sessions.Create(ctx, &session_proto2.SessionId{Id: int32(newUser.Id)})
	//if err != nil {
	//	return &userProto.UserResponse{
	//		Code: 500,
	//		User: nil,
	//	}, err
	//}

	return &userProto.UserResponse{
		User:  u.Usecase.User2ProtoUser(newUser),
		Token: "",
	}, nil
}

func (u UserServer) DeleteUser(ctx context.Context, nothing *userProto.UserNothing) (*userProto.UserNothing, error) {
	cookieId, ok := u.Usecase.GetParamFromContext(ctx, "cookieId")
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		return &userProto.UserNothing{}, status.Error(403, response.Error())
	}

	urlId, ok := u.Usecase.GetParamFromContext(ctx, "urlId")
	if !ok {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["id"] = "Юзера с таким id нет"
		return &userProto.UserNothing{}, status.Error(403, response.Error())
	}

	log.Print(cookieId)
	log.Print(urlId)

	if cookieId != urlId {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Отказано в доступе"}
		response.Description["id"] = "Пытаешься удалить не себя"
		return &userProto.UserNothing{}, status.Error(403, response.Error())
	}

	err := u.Usecase.DeleteUser(cookieId)
	if err != nil {
		return &userProto.UserNothing{}, status.Error(500, "")
	}

	_, err = u.Sessions.Delete(ctx, &session_proto2.SessionId{Id: int32(cookieId)})
	if err != nil {
		return &userProto.UserNothing{}, status.Error(500, "")
	}

	return &userProto.UserNothing{}, nil
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

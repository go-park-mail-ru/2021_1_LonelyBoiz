package delivery

import (
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	session_proto2 "server/internal/auth_server/delivery/session"
	chatUsecase "server/internal/pkg/chat/usecase"
	messageUsecase "server/internal/pkg/message/usecase"
	model "server/internal/pkg/models"
	"server/internal/pkg/user/usecase"
	userProto "server/internal/user_server/delivery/proto"
	"strconv"
)

type UserServer struct {
	userProto.UnimplementedUserServiceServer
	UserUsecase    usecase.UserUseCaseInterface
	ChatUsecase    chatUsecase.ChatUsecaseInterface
	MessageUsecase messageUsecase.MessageUsecaseInterface
	Sessions       session_proto2.AuthCheckerClient
}

func (u UserServer) CreateUser(ctx context.Context, user *userProto.User) (*userProto.UserResponse, error) {
	nUser, ok := u.UserUsecase.ProtoUser2User(user)
	if !ok {
		return &userProto.UserResponse{}, status.Error(500, "")
	}

	newUser, code, responseError := u.UserUsecase.CreateNewUser(nUser)
	if code != 200 {
		return &userProto.UserResponse{}, status.Error(codes.Code(code), responseError.Error())
	}

	token, err := u.Sessions.Create(ctx, &session_proto2.SessionId{Id: int32(newUser.Id)})
	if err != nil {
		return &userProto.UserResponse{}, status.Error(codes.Code(code), err.Error())
	}

	protoUser, ok := u.UserUsecase.User2ProtoUser(newUser)
	if !ok {
		return &userProto.UserResponse{}, status.Error(500, "")
	}

	return &userProto.UserResponse{
		User:  protoUser,
		Token: token.GetToken(),
	}, nil
}

func (u UserServer) DeleteUser(ctx context.Context, nothing *userProto.UserNothing) (*userProto.UserNothing, error) {
	id, code, err := u.UserUsecase.CheckIds(ctx)
	if err != nil {
		return &userProto.UserNothing{}, status.Error(codes.Code(code), err.Error())
	}

	err = u.UserUsecase.DeleteUser(id)
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
	id, code, err := u.UserUsecase.CheckIds(ctx)
	if err != nil {
		return &userProto.User{}, status.Error(codes.Code(code), err.Error())
	}

	nUser, ok := u.UserUsecase.ProtoUser2User(user)
	if !ok {
		return &userProto.User{}, status.Error(500, "")
	}

	newUser, code, err := u.UserUsecase.ChangeUserInfo(nUser, id)
	if code != 200 {
		if err != nil {
			return &userProto.User{}, status.Error(codes.Code(code), err.Error())
		}
		return &userProto.User{}, status.Error(codes.Code(code), "")
	}

	protoUser, ok := u.UserUsecase.User2ProtoUser(newUser)
	if !ok {
		return &userProto.User{}, status.Error(500, "")
	}

	if protoUser == nil {
		return &userProto.User{}, status.Error(500, "")
	}

	return protoUser, nil
}

func (u UserServer) CheckUser(ctx context.Context, login *userProto.UserLogin) (*userProto.UserResponse, error) {
	newUser, code, err := u.UserUsecase.SignIn(model.User{
		Email:          login.GetEmail(),
		Password:       login.GetPassword(),
		SecondPassword: login.GetSecondPassword(),
	})

	if code == 500 {
		return &userProto.UserResponse{}, status.Error(codes.Code(code), "")
	}

	if code != 200 {
		return &userProto.UserResponse{}, status.Error(codes.Code(code), err.Error())
	}

	token, err := u.Sessions.Create(ctx, &session_proto2.SessionId{Id: int32(newUser.Id)})
	if err != nil {
		return &userProto.UserResponse{}, status.Error(500, "")
	}

	protoUser, ok := u.UserUsecase.User2ProtoUser(newUser)
	if !ok {
		return &userProto.UserResponse{}, status.Error(500, "")
	}

	return &userProto.UserResponse{User: protoUser, Token: token.GetToken()}, nil
}

func (u UserServer) GetUserById(ctx context.Context, nothing *userProto.UserNothing) (*userProto.User, error) {
	id, ok := u.UserUsecase.GetParamFromContext(ctx, "urlId")
	if !ok {
		id, ok = u.UserUsecase.GetParamFromContext(ctx, "cookieId")

		if !ok {
			response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
			return &userProto.User{}, status.Error(403, response.Error())
		}
	}
	userInfo, err := u.UserUsecase.GetUserInfoById(id)
	if err != nil {
		response := model.ErrorResponse{Err: "Пользователь не найден"}
		return &userProto.User{}, status.Error(401, response.Error())
	}

	protoUser, ok := u.UserUsecase.User2ProtoUser(userInfo)
	if !ok {
		return &userProto.User{}, status.Error(500, "")
	}

	return protoUser, nil
}

func (u UserServer) CreateFeed(ctx context.Context, nothing *userProto.UserNothing) (*userProto.Feed, error) {
	id, ok := u.UserUsecase.GetParamFromContext(ctx, "cookieId")
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		return &userProto.Feed{}, status.Error(403, response.Err)
	}

	limitInt, ok := u.UserUsecase.GetParamFromContext(ctx, "urlCount")
	if !ok {
		response := model.ErrorResponse{Err: "Неверный формат count"}
		return &userProto.Feed{}, status.Error(400, response.Err)
	}

	feed, code, err := u.UserUsecase.CreateFeed(id, limitInt)
	if code == 500 {
		return &userProto.Feed{}, status.Error(500, "")
	}
	if code != 200 {
		return &userProto.Feed{}, status.Error(codes.Code(code), err.Error())
	}

	var feed1 []*userProto.UserId

	for _, idFromFeed := range feed {
		feed1 = append(feed1, &userProto.UserId{Id: int32(idFromFeed)})
	}

	return &userProto.Feed{Users: feed1}, nil
}

func (u UserServer) CreateChat(ctx context.Context, like *userProto.Like) (*userProto.Chat, error) {
	userId, ok := u.UserUsecase.GetParamFromContext(ctx, "cookieId")
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		return &userProto.Chat{}, status.Error(403, response.Err)
	}

	chat, code, err := u.UserUsecase.CreateChat(userId, model.Like{
		UserId:   int(like.GetUserId()),
		Reaction: like.GetReaction(),
	})

	if code == 204 {
		return &userProto.Chat{}, status.Error(204, "Success like")
	}

	if code == 500 {
		return &userProto.Chat{}, status.Error(500, "")
	}

	if err != nil {
		return &userProto.Chat{}, status.Error(codes.Code(code), err.Error())
	}

	photos, ok := u.UserUsecase.Photos2ProtoPhotos(chat.Photos)
	if !ok {
		return &userProto.Chat{}, status.Error(500, "")
	}

	return &userProto.Chat{
		ChatId:              int32(chat.ChatId),
		PartnerId:           int32(chat.PartnerId),
		PartnerName:         chat.PartnerName,
		LastMessage:         chat.LastMessage,
		LastMessageTime:     chat.LastMessageTime,
		LastMessageAuthorId: int32(chat.LastMessageAuthorId),
		Photos:              photos,
	}, nil
}

func (u UserServer) GetChats(ctx context.Context, nothing *userProto.UserNothing) (*userProto.ChatsResponse, error) {
	userId, ok := u.UserUsecase.GetParamFromContext(ctx, "cookieId")
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		return &userProto.ChatsResponse{}, status.Error(403, response.Err)
	}

	limit, ok := u.UserUsecase.GetParamFromContext(ctx, "urlCount")
	if !ok {
		response := model.ErrorResponse{Err: "Неверный формат count"}
		return &userProto.ChatsResponse{}, status.Error(400, response.Err)
	}

	offset, ok := u.UserUsecase.GetParamFromContext(ctx, "urlOffset")
	if !ok {
		response := model.ErrorResponse{Err: "Неверный формат offset"}
		return &userProto.ChatsResponse{}, status.Error(400, response.Err)
	}

	chats, err := u.ChatUsecase.GetChat(userId, limit, offset)
	if err != nil {
		return &userProto.ChatsResponse{}, status.Error(500, "")
	}

	//TODO: пункт 1
	return chats, nil
}

func (u UserServer) GetMessages(ctx context.Context, nothing *userProto.UserNothing) (*userProto.MessagesResponse, error) {
	userId, ok := u.UserUsecase.GetParamFromContext(ctx, "cookieId")
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		return &userProto.MessagesResponse{}, status.Error(403, response.Err)
	}

	limit, ok := u.UserUsecase.GetParamFromContext(ctx, "urlCount")
	if !ok {
		response := model.ErrorResponse{Err: "Неверный формат count"}
		return &userProto.MessagesResponse{}, status.Error(400, response.Err)
	}

	offset, ok := u.UserUsecase.GetParamFromContext(ctx, "urlOffset")
	if !ok {
		response := model.ErrorResponse{Err: "Неверный формат offset"}
		return &userProto.MessagesResponse{}, status.Error(400, response.Err)
	}

	chatId, ok := u.UserUsecase.GetParamFromContext(ctx, "urlChatId")
	if !ok {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["messageId"] = "Сообщения с таким id нет"
		return &userProto.MessagesResponse{}, status.Error(400, response.Error())
	}

	messages, code, err := u.MessageUsecase.ManageMessage(userId, chatId, limit, offset)

	if code != 200 {
		return &userProto.MessagesResponse{}, status.Error(codes.Code(code), err.Error())
	}
	// TODO::
	return messages, nil
}

func (u UserServer) CreateMessage(ctx context.Context, message *userProto.Message) (*userProto.Message, error) {
	id, ok := u.UserUsecase.GetParamFromContext(ctx, "cookieId")
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		return &userProto.Message{}, status.Error(403, response.Err)
	}

	chatId, ok := u.UserUsecase.GetParamFromContext(ctx, "urlChatId")
	if !ok {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["messageId"] = "Сообщения с таким id нет"
		return &userProto.Message{}, status.Error(400, response.Error())
	}

	newMessage, code, err := u.MessageUsecase.CreateMessage(model.Message{
		MessageId:    int(message.GetMessageId()),
		AuthorId:     int(message.GetAuthorId()),
		ChatId:       int(message.GetChatId()),
		Text:         message.GetText(),
		Reaction:     int(message.GetReaction()),
		Time:         message.GetTime(),
		MessageOrder: int(message.GetMessageOrder()),
	}, chatId, id)

	if code != 200 {
		return &userProto.MessagesResponse{}, status.Error(codes.Code(code), err.Error())
	}

	// TODO::
	return newMessage, nil
}

func (u UserServer) ChangeMessage(ctx context.Context, message *userProto.Message) (*userProto.Message, error) {
	userId, ok := u.UserUsecase.GetParamFromContext(ctx, "cookieId")
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		return &userProto.Message{}, status.Error(403, response.Err)
	}

	messageId, ok := u.UserUsecase.GetParamFromContext(ctx, "urlMessageId")
	if !ok {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["messageId"] = "Сообщения с таким id нет"
		return &userProto.Message{}, status.Error(400, response.Error())
	}

	newMessage, code, err := u.MessageUsecase.ChangeMessage(userId, messageId, message)
	if code != 204 {
		return &userProto.Message{}, status.Error(codes.Code(code), err.Error())
	}

	return newMessage, nil
}

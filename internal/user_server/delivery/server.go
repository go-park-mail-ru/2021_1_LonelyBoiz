package delivery

import (
	session_proto2 "server/internal/auth_server/delivery/session"
	chatUsecase "server/internal/pkg/chat/usecase"
	messageUsecase "server/internal/pkg/message/usecase"
	model "server/internal/pkg/models"
	"server/internal/pkg/user/usecase"
	userProto "server/internal/user_server/delivery/proto"
	"strconv"

	"golang.org/x/net/context"
	codes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	userProto.UnimplementedUserServiceServer
	UserUsecase    usecase.UserUseCaseInterface
	ChatUsecase    chatUsecase.ChatUsecaseInterface
	MessageUsecase messageUsecase.MessageUsecaseInterface
	Sessions       session_proto2.AuthCheckerClient
}

func (u UserServer) DeleteChat(ctx context.Context, user *userProto.UserNothing) (*userProto.UserNothing, error) {
	chatId, ok := u.UserUsecase.GetParamFromContext(ctx, "urlChatId")
	if !ok {
		response := model.ErrorDescriptionResponse{Description: map[string]string{}, Err: "Неверный формат входных данных"}
		response.Description["chatId"] = "Чата с таким id нет"
		return &userProto.UserNothing{}, status.Error(400, response.Err)
	}

	err := u.UserUsecase.DeleteChat(chatId)
	if err != nil {
		return &userProto.UserNothing{}, status.Error(500, "")
	}

	return &userProto.UserNothing{}, nil
}

func (u UserServer) CreateUser(ctx context.Context, user *userProto.User) (*userProto.UserResponse, error) {
	newUser, code, responseError := u.UserUsecase.CreateNewUser(u.UserUsecase.ProtoUser2User(user))
	if code != 200 {
		return &userProto.UserResponse{}, status.New(codes.Code(code), responseError.Error()).Err()
	}
	u.UserUsecase.LogInfo("Аккаунт с id = " + strconv.Itoa(newUser.Id) + " создан")

	token, err := u.Sessions.Create(ctx, &session_proto2.SessionId{Id: int32(newUser.Id)})
	if err != nil {
		u.UserUsecase.LogError(err.Error())
		return &userProto.UserResponse{}, status.Error(codes.Code(code), err.Error())
	}
	u.UserUsecase.LogInfo("Для аккаунта с id = " + strconv.Itoa(newUser.Id) + " создан токен")

	return &userProto.UserResponse{
		User:  u.UserUsecase.User2ProtoUser(newUser),
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

	newUser, code, err := u.UserUsecase.ChangeUserInfo(u.UserUsecase.ProtoUser2User(user), id)
	if code != 200 {
		if err != nil {
			return &userProto.User{}, status.Error(codes.Code(code), err.Error())
		}
		return &userProto.User{}, status.Error(codes.Code(code), "")
	}

	return u.UserUsecase.User2ProtoUser(newUser), nil
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

	return &userProto.UserResponse{User: u.UserUsecase.User2ProtoUser(newUser), Token: token.GetToken()}, nil
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

	return u.UserUsecase.User2ProtoUser(userInfo), nil
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

	var protoFeed []*userProto.UserId

	for _, idFromFeed := range feed {
		protoFeed = append(protoFeed, &userProto.UserId{Id: int32(idFromFeed)})
	}

	return &userProto.Feed{Users: protoFeed}, nil
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

	return &userProto.Chat{
		ChatId:              int32(chat.ChatId),
		PartnerId:           int32(chat.PartnerId),
		PartnerName:         chat.PartnerName,
		LastMessage:         chat.LastMessage,
		LastMessageTime:     chat.LastMessageTime,
		LastMessageAuthorId: int32(chat.LastMessageAuthorId),
		Photos:              u.UserUsecase.Photos2ProtoPhotos(chat.Photos),
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

	var protoChats []*userProto.Chat
	for _, chat := range chats {
		protoChats = append(protoChats, u.ChatUsecase.Chat2ProtoChat(chat))
	}

	return &userProto.ChatsResponse{Chats: protoChats}, nil
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

	var protoMessages []*userProto.Message
	for _, message := range messages {
		protoMessages = append(protoMessages, u.MessageUsecase.Message2ProtoMessage(message))
	}

	return &userProto.MessagesResponse{Messages: protoMessages}, nil
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
		response.Description["chatId"] = "Пользователь вас заблокирвоал"
		return &userProto.Message{}, status.Error(400, response.Error())
	}

	newMessage, code, err := u.MessageUsecase.CreateMessage(u.MessageUsecase.ProtoMessage2Message(message), chatId, id)
	if code != 200 {
		return &userProto.Message{}, status.Error(codes.Code(code), err.Error())
	}

	return u.MessageUsecase.Message2ProtoMessage(newMessage), nil
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

	newMessage, code, err := u.MessageUsecase.ChangeMessage(userId, messageId, u.MessageUsecase.ProtoMessage2Message(message))
	if code != 204 {
		return &userProto.Message{}, status.Error(codes.Code(code), err.Error())
	}

	return u.MessageUsecase.Message2ProtoMessage(newMessage), nil
}

func (u UserServer) AddToSecreteAlbum(ctx context.Context, message *userProto.User) (*userProto.UserNothing, error) {
	ownerId, ok := u.UserUsecase.GetParamFromContext(ctx, "cookieId")
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		return &userProto.UserNothing{}, status.Error(403, response.Err)
	}

	code, err := u.UserUsecase.AddToSecreteAlbum(ownerId, u.UserUsecase.ProtoPhotos2Photos(message.Photos))
	if err != nil {
		return &userProto.UserNothing{}, status.Error(codes.Code(code), err.Error())
	}

	return &userProto.UserNothing{}, nil
}

func (u UserServer) UnlockSecretAlbum(ctx context.Context, message *userProto.UserNothing) (*userProto.UserNothing, error) {
	ownerId, ok := u.UserUsecase.GetParamFromContext(ctx, "cookieId")
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		return &userProto.UserNothing{}, status.Error(403, response.Err)
	}

	getterId, ok := u.UserUsecase.GetParamFromContext(ctx, "getterId")
	if !ok {
		response := model.ErrorResponse{Err: "Пользователь не найден"}
		return &userProto.UserNothing{}, status.Error(400, response.Error())
	}

	code, err := u.UserUsecase.UnblockSecreteAlbum(ownerId, getterId)
	if err != nil {
		return &userProto.UserNothing{}, status.Error(codes.Code(code), err.Error())
	}

	return &userProto.UserNothing{}, nil
}

func (u UserServer) GetSecreteAlbum(ctx context.Context, message *userProto.UserNothing) (*userProto.Photos, error) {
	ownerId, ok := u.UserUsecase.GetParamFromContext(ctx, "ownerId")
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		return &userProto.Photos{}, status.Error(403, response.Err)
	}

	getterId, ok := u.UserUsecase.GetParamFromContext(ctx, "cookieId")
	if !ok {
		response := model.ErrorResponse{Err: "Пользователь не найден"}
		return &userProto.Photos{}, status.Error(400, response.Error())
	}

	photos, code, err := u.UserUsecase.GetSecreteAlbum(ownerId, getterId)
	if err != nil {
		return &userProto.Photos{}, status.Error(codes.Code(code), err.Error())
	}

	protoPhotos := userProto.Photos{
		Photos: u.UserUsecase.Photos2ProtoPhotos(photos),
	}

	return &protoPhotos, nil
}

func (u UserServer) BlockSecretAlbum(ctx context.Context, message *userProto.UserNothing) (*userProto.UserNothing, error) {
	ownerId, ok := u.UserUsecase.GetParamFromContext(ctx, "cookieId")
	if !ok {
		response := model.ErrorResponse{Err: model.SessionErrorDenAccess}
		return &userProto.UserNothing{}, status.Error(403, response.Err)
	}

	getterId, ok := u.UserUsecase.GetParamFromContext(ctx, "getterId")
	if !ok {
		response := model.ErrorResponse{Err: "Пользователь не найден"}
		return &userProto.UserNothing{}, status.Error(400, response.Error())
	}

	code, err := u.UserUsecase.BlockSecreteAlbum(ownerId, getterId)
	if err != nil {
		return &userProto.UserNothing{}, status.Error(codes.Code(code), err.Error())
	}

	return &userProto.UserNothing{}, nil
}

package delivery

import (
	"errors"
	session_proto "server/internal/auth_server/delivery/session"
	auth_server "server/internal/pkg/session"

	"golang.org/x/net/context"
)

type AuthServer struct {
	Usecase auth_server.SessionManagerInterface
	session_proto.UnimplementedAuthCheckerServer
}

// Create создает токен, который необходимо вставить в заголовок Set-Cookie
func (a AuthServer) Create(ctx context.Context, id *session_proto.SessionId) (*session_proto.SessionToken, error) {
	token, err := a.Usecase.SetSession(int(id.GetId()))

	if err != nil {
		return nil, err
	}

	return &session_proto.SessionToken{Token: token}, nil
}

// Check проверят куку
func (a AuthServer) Check(ctx context.Context, token *session_proto.SessionToken) (*session_proto.SessionId, error) {
	id, ok := a.Usecase.CheckSession([]string{token.GetToken()})
	if !ok {
		return nil, errors.New("пользователь не найден")
	}

	return &session_proto.SessionId{Id: int32(id)}, nil
}

// Delete вызывается для удаление кук, при удалении аккаунта
func (a AuthServer) Delete(ctx context.Context, id *session_proto.SessionId) (*session_proto.Nothing, error) {
	_ = a.Usecase.DeleteSessionById(int(id.GetId()))
	return &session_proto.Nothing{Dummy: true}, nil
}

package delivery

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"server/internal/auth_server/delivery/session"
	auth_server "server/internal/pkg/session"
)

type AuthServer struct {
	Usecase auth_server.SessionManagerInterface
	session_proto.UnimplementedAuthCheckerServer
}

// Create создает токен, который необходимо вставить в заголовок Set-Cookie
func (a AuthServer) Create(ctx context.Context, id *session_proto.SessionId) (*session_proto.SessionToken, error) {
	token, err := a.Usecase.SetSession(int(id.GetId()))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(token)

	return &session_proto.SessionToken{Token: token}, nil
}

// Check проверят куку
func (a AuthServer) Check(ctx context.Context, token *session_proto.SessionToken) (*session_proto.SessionId, error) {
	fmt.Println(token.GetToken())
	id, ok := a.Usecase.CheckSession([]string{token.GetToken()})
	if ok == false {
		return nil, errors.New("user Not Found")
	}
	fmt.Println(id)
	return &session_proto.SessionId{Id: int32(id)}, nil
}

// Delete вызывается для удаление кук, при удалении аккаунта
func (a AuthServer) Delete(ctx context.Context, id *session_proto.SessionId) (*session_proto.Nothing, error) {
	_ = a.Usecase.DeleteSessionById(int(id.GetId()))
	return &session_proto.Nothing{Dummy: true}, nil
}

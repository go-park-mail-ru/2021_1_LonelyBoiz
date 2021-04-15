package delivery

import (
	"net/http"
	"server/internal/pkg/session"
	"server/internal/pkg/user/usecase"
)

type UserHandlerInterface interface {
	UploadPhoto(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	LogOut(w http.ResponseWriter, r *http.Request)
	ChangeUserInfo(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	GetUserInfo(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetLogin(w http.ResponseWriter, r *http.Request)
	WsHandler(w http.ResponseWriter, r *http.Request)
	LikesHandler(w http.ResponseWriter, r *http.Request)
	WebSocketChatResponse()
	DownloadPhoto(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	UserCase usecase.UserUsecaseInterface
	Sessions *session.SessionsManager
}

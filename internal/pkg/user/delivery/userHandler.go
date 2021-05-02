package delivery

import (
	"github.com/gorilla/mux"
	"net/http"
	session_proto2 "server/internal/auth_server/delivery/session"
	"server/internal/pkg/user/usecase"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserCase usecase.UserUseCaseInterface
	Sessions session_proto2.AuthCheckerClient
}

type UserHandlerInterface interface {
	// остается
	SignIn(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	//LogOut(w http.ResponseWriter, r *http.Request)
	ChangeUserInfo(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	GetUserInfo(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetLogin(w http.ResponseWriter, r *http.Request)

	WsHandler(w http.ResponseWriter, r *http.Request)
	LikesHandler(w http.ResponseWriter, r *http.Request)
}

func (a *UserHandler) SetHandlersWithCheckCookie(subRouter *mux.Router) {
	// получить ленту
	subRouter.HandleFunc("/feed", a.GetUsers).Methods("GET")
	// првоерить куку
	subRouter.HandleFunc("/auth", a.GetLogin).Methods("GET")
	// удалить юзера
	subRouter.HandleFunc("/users/{id:[0-9]+}", a.DeleteUser).Methods("DELETE")
	// изменить информацию о юзере
	subRouter.HandleFunc("/users/{id:[0-9]+}", a.ChangeUserInfo).Methods("PATCH")
	// поставить оценку юзеру из ленты
	subRouter.HandleFunc("/likes", a.LikesHandler).Methods("POST")
	// открытие вэбсокетного соединения
	subRouter.HandleFunc("/ws", a.WsHandler).Methods("GET")
}

func (a *UserHandler) SetHandlersWithoutCheckCookie(subRouter *mux.Router) {
	subRouter.HandleFunc("/users/{id:[0-9]+}", a.GetUserInfo).Methods("GET")
	// регистрация
	subRouter.HandleFunc("/users", a.SignUp).Methods("POST")
	// логин
	subRouter.HandleFunc("/login", a.SignIn).Methods("POST")
}

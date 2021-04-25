package delivery

import (
	"net/http"
	"server/internal/pkg/session"
	"server/internal/pkg/user/usecase"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserCase usecase.UserUseCaseInterface
	Sessions session.SessionManagerInterface
}

type UserHandlerInterface interface {
	// остается
	SignIn(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	LogOut(w http.ResponseWriter, r *http.Request)
	ChangeUserInfo(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	GetUserInfo(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetLogin(w http.ResponseWriter, r *http.Request)

	UnblockSecreteAlbum(w http.ResponseWriter, r *http.Request)
	GetSecreteAlbum(w http.ResponseWriter, r *http.Request)

	WsHandler(w http.ResponseWriter, r *http.Request)
	LikesHandler(w http.ResponseWriter, r *http.Request)
}

func (a *UserHandler) SetRawRouter(subRouter *mux.Router) {
	// оплата
	subRouter.HandleFunc("/pay", a.Payment).Methods("POST")
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

	// добавить фотки в секретный альбом
	subRouter.HandleFunc("/secretAlbum", a.AddToSecreteAlbum).Methods("POST")
	// разблокировать секретный альбом
	subRouter.HandleFunc("/unlockSecretAlbum/{getterId:[0-9]+}", a.UnblockSecreteAlbum).Methods("POST")
	// посмотреть секретный альбом
	subRouter.HandleFunc("/secretAlbum/{ownerId:[0-9]+}", a.GetSecreteAlbum).Methods("Get")
	// открытие вэбсокетного соединения
	subRouter.HandleFunc("/ws", a.WsHandler).Methods("GET")
}

func (a *UserHandler) SetHandlersWithoutCheckCookie(subRouter *mux.Router) {
	subRouter.HandleFunc("/users/{id:[0-9]+}", a.GetUserInfo).Methods("GET")
	// регистрация
	subRouter.HandleFunc("/users", a.SignUp).Methods("POST")
	// логин
	subRouter.HandleFunc("/login", a.SignIn).Methods("POST")
	// логаут
	subRouter.HandleFunc("/login", a.LogOut).Methods("DELETE")
}

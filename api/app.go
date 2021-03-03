package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Email    string
	Password string
}

type App struct {
	addr   string
	router *mux.Router
	Users  []User
}

func (a *App) Start() error {
	fmt.Println("server start")
	err := http.ListenAndServe(a.addr, a.router)
	if err != nil {
		return err
	}

	return nil
}

type Config struct {
	addr   string
	router *mux.Router
}

func NewConfig() Config {
	newConfig := Config{}
	newConfig.addr = ":8000"
	newConfig.router = mux.NewRouter()
	return newConfig
}

func (a *App) InitializeRoutes(currConfig Config) {
	a.addr = currConfig.addr
	a.router = currConfig.router
	a.router.HandleFunc("/login", a.SignIn).Methods("POST")
	a.router.HandleFunc("/users", a.SignUp).Methods("POST")
	a.router.HandleFunc("/users/{id:[0-9]+}", a.ChangeUserInfo).Methods("PATCH")
	a.router.HandleFunc("/users/{id:[0-9]+}", a.GetUserInfo).Methods("GET")
	a.router.HandleFunc("/users/{id:[0-9]+}/photos", a.UploadPhoto).Methods("POST")
	a.router.HandleFunc("/users/{id:[0-9]+}/photos/{id:[0-9]+}", a.DownloadPhoto).Methods("GET")
	a.router.HandleFunc("/users/{id:[0-9]+}/photos/{id:[0-9]+}", a.DeletePhoto).Methods("DELETE")
}

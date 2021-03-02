package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Email    string
	Password string
}

type App struct {
	Router *mux.Router
	Users  []User
}

func (a *App) Run(addr string) {
	fmt.Println("server start")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) InitializeRoutes() {
	a.Router.HandleFunc("/login", a.SignIn).Methods("POST")
	a.Router.HandleFunc("/users", a.SignUp).Methods("POST")
	a.Router.HandleFunc("/users/{id:[0-9]+}", a.ChangeUserInfo).Methods("POST")
	a.Router.HandleFunc("/users/{id:[0-9]+}", a.GetUserInfo).Methods("GET")
	a.Router.HandleFunc("/users/{id:[0-9]+}/photos", a.UploadPhoto).Methods("POST")
	a.Router.HandleFunc("/users/{id:[0-9]+}/photos/{id:[0-9]+}", a.DownloadPhoto).Methods("GET")
	a.Router.HandleFunc("/users/{id:[0-9]+}/photos/{id:[0-9]+}", a.DeletePhoto).Methods("DELETE")
}

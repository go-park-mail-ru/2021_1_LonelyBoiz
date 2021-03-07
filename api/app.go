package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const charSet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789"

func KeyGen() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 40)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(b)
}

type User struct {
	Id             int
	Email          string    `json:"mail"`
	Password       string    `json:"pass,omitempty"`
	SecondPassword string    `json:"passRepeat,omitempty"`
	PasswordHash   []byte    `json:",omitempty"`
	OldPassword    string    `json:"oldPass,omitempty"`
	Name           string    `json:"name"`
	Birthday       time.Time `json:"birthday"`
	Description    string    `json:"description"`
	City           string    `json:"city"`
	AvatarAddr     string    `json:"avatar"`
	Instagram      string    `json:"instagram"`
	Sex            string    `json:"sex"`
	DatePreference string    `json:"datePreference"`
}

type App struct {
	addr     string
	router   *mux.Router
	Users    []User
	UserIds  int
	Sessions map[int][]http.Cookie
}

type errorResponse struct {
	Description map[string]string
	Err         string
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
	addr    string
	userIds int
	router  *mux.Router
}

func NewConfig() Config {
	newConfig := Config{}
	newConfig.addr = ":8003"
	newConfig.userIds = 0
	newConfig.router = mux.NewRouter()
	return newConfig
}

func (a *App) InitializeRoutes(currConfig Config) {
	a.addr = currConfig.addr
	a.router = currConfig.router
	a.UserIds = currConfig.userIds
	a.Sessions = make(map[int][]http.Cookie)
	a.router.HandleFunc("/login", a.SignIn).Methods("POST")
	a.router.HandleFunc("/users", a.SignUp).Methods("POST")
	a.router.HandleFunc("/users/{id:[0-9]+}", a.ChangeUserInfo).Methods("PATCH")
	a.router.HandleFunc("/users/{id:[0-9]+}", a.GetUserInfo).Methods("GET")
	a.router.HandleFunc("/users/{id:[0-9]+}/photos", a.UploadPhoto).Methods("POST")
	a.router.HandleFunc("/users/{id:[0-9]+}/photos/{id:[0-9]+}", a.DownloadPhoto).Methods("GET")
	a.router.HandleFunc("/users/{id:[0-9]+}/photos/{id:[0-9]+}", a.DeletePhoto).Methods("DELETE")
	a.router.HandleFunc("/users/{id:[0-9]+}", a.DeleteUser).Methods("DELETE")
	a.router.HandleFunc("/login", a.LogOut).Methods("DELETE")
}

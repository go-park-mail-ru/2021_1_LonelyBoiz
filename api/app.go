package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const charSet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789"

func KeyGen() string {
	b := make([]byte, 40)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(b)
}

type User struct {
	Id             int
	Email          string   `json:"mail"`
	Password       string   `json:"pass,omitempty"`
	SecondPassword string   `json:"passRepeat,omitempty"`
	PasswordHash   []byte   `json:",omitempty"`
	OldPassword    string   `json:"oldPass,omitempty"`
	Name           string   `json:"name"`
	Birthday       int64    `json:"birthday"`
	Description    string   `json:"description"`
	City           string   `json:"city"`
	AvatarAddr     []string `json:"avatar"`
	Instagram      string   `json:"instagram"`
	Sex            string   `json:"sex"`
	DatePreference string   `json:"datePreference"`
}

type App struct {
	addr     string
	router   *mux.Router
	Users    map[int]User
	UserIds  int
	Sessions map[int][]http.Cookie
}

func (a *App) Start() error {
	fmt.Println("Server start")

	ch := handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"}),
		handlers.AllowCredentials())

	s := http.Server{
		Addr:         a.addr,
		Handler:      ch(a.router),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err := s.ListenAndServe()
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
	newConfig.addr = ":8001"
	newConfig.userIds = 0
	newConfig.router = mux.NewRouter()
	return newConfig
}

func (a *App) InitializeRoutes(currConfig Config) {
	rand.Seed(time.Now().UnixNano())

	a.addr = currConfig.addr
	a.router = currConfig.router
	a.UserIds = currConfig.userIds
	a.Sessions = make(map[int][]http.Cookie)
	a.Users = make(map[int]User)

	a.router.HandleFunc("/login", a.SignIn).Methods("POST")
	a.router.HandleFunc("/users", a.SignUp).Methods("POST")
	a.router.HandleFunc("/users/{id:[0-9]+}", a.ChangeUserInfo).Methods("PATCH")
	a.router.HandleFunc("/users/{id:[0-9]+}", a.GetUserInfo).Methods("GET")
	a.router.HandleFunc("/users/{id:[0-9]+}/photos", a.UploadPhoto).Methods("POST")
	a.router.HandleFunc("/users/{id:[0-9]+}/photos", a.DeletePhoto).Methods("DELETE")
	a.router.HandleFunc("/users/{id:[0-9]+}", a.DeleteUser).Methods("DELETE")
	a.router.HandleFunc("/login", a.LogOut).Methods("DELETE")
}

/*curl -H "Origin: http://localhost:3000" \
-H "Access-Control-Request-Method: POST" \
-H "Access-Control-Request-Headers: X-Requested-With" \
-X OPTIONS --verbose http://localhost/8001/login
*/

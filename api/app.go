package api

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"math/rand"
	"net/http"
	"os"
	model "server/models"
	"server/repository"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const charSet = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789"

func init() {
	govalidator.CustomTypeTagMap.Set(
		"ageValid",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			birthday, ok := i.(int64)
			if !ok {
				return false
			}

			tm := time.Unix(birthday, 0)
			diff := time.Now().Sub(tm)

			if diff/24/365 < 18 {
				return false
			}
			return true
		}),
	)
}

func KeyGen() string {
	b := make([]byte, 40)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(b)
}

type App struct {
	addr     string
	router   *mux.Router
	Users    map[int]model.User
	UserIds  int
	Sessions map[int][]http.Cookie
	mutex    *sync.Mutex
	Db       repository.RepoSqlx
}

func (a *App) Start() error {
	fmt.Println("Server start")

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://lepick.herokuapp.com"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With"},
		Debug:            true,
	})

	corsHandler := cors.Handler(a.router)

	s := http.Server{
		Addr:         a.addr,
		Handler:      corsHandler,
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
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	newConfig.addr = ":" + port
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
	a.Users = make(map[int]model.User)
	a.Db = repository.Init()

	// валидация кук
	subRouter := a.router.NewRoute().Subrouter()
	subRouter.Use(a.MiddlewareValidateCookie)

	subRouter.HandleFunc("/users/{id:[0-9]+}", a.GetUserInfo).Methods("GET")
	//subRouter.HandleFunc("/users", a.GetUsers).Methods("GET")
	subRouter.HandleFunc("/auth", a.GetLogin).Methods("GET")
	subRouter.HandleFunc("/users/{id:[0-9]+}", a.DeleteUser).Methods("DELETE")
	subRouter.HandleFunc("/users/{id:[0-9]+}", a.ChangeUserInfo).Methods("PATCH")

	// валидация всех данных, без кук
	a.router.HandleFunc("/users", a.SignUp).Methods("POST")

	// валидация пароля
	a.router.HandleFunc("/login", a.SignIn).Methods("POST")

	// не требуется
	a.router.HandleFunc("/login", a.LogOut).Methods("DELETE")

	//a.router.HandleFunc("/users/{id:[0-9]+}/photos", a.UploadPhoto).Methods("POST")
	//a.router.HandleFunc("/users/{id:[0-9]+}/photos", a.DeletePhoto).Methods("DELETE")
}

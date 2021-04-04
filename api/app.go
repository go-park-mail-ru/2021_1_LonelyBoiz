package api

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"math/rand"
	"net/http"
	"os"
	"server/repository"
	"time"

	"github.com/gorilla/mux"
	cors2 "github.com/rs/cors"
	"github.com/sirupsen/logrus"
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
	addr   string
	router *mux.Router
	Db     repository.RepoSqlx
	Logger *logrus.Entry
}

func (a *App) Start() error {
	fmt.Println("Server start")

	cors := cors2.New(cors2.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://lepick.herokuapp.com"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With"},
		Debug:            false,
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
	//a.UserIds = currConfig.userIds
	a.Db = repository.Init()

	contextLogger := logrus.WithFields(logrus.Fields{
		"mode": "[access_log]",
	})
	logrus.SetFormatter(&logrus.TextFormatter{})
	a.Logger = contextLogger

	a.router.Use(a.MiddlewareLogger)
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

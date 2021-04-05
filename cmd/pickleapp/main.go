package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	cors2 "github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"os"
	"server/internal/pickleapp/middleware"
	repository2 "server/internal/pickleapp/repository"
	"server/internal/pkg/session"
	repository3 "server/internal/pkg/session/repository"
	handler "server/internal/pkg/user/delivery"
	"server/internal/pkg/user/repository"
	"server/internal/pkg/user/usecase"
	"time"
)

type App struct {
	addr   string
	router *mux.Router
	Db     *sqlx.DB
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

	//конфиг
	a.addr = currConfig.addr
	a.router = currConfig.router

	// логгер
	contextLogger := logrus.WithFields(logrus.Fields{
		"mode": "[access_log]",
	})
	logrus.SetFormatter(&logrus.TextFormatter{})
	a.Logger = contextLogger

	// бд
	a.Db = repository2.Init()
	userRep := repository.UserRepository{DB: a.Db}
	sessionRep := repository3.SessionRepository{DB: a.Db}

	userUcase := usecase.UserUsecase{Db: userRep}
	sessionManager := session.SessionsManager{DB: sessionRep}

	userHandler := handler.UserHandler{
		Db:       userRep,
		Logger:   a.Logger,
		UserCase: userUcase,
		Sessions: &sessionManager,
	}

	checkcookiem := middleware.ValidateCookieMiddleware{Session: &sessionManager}

	//a.router.Use(middleware.MiddlewareLogger)
	// валидация кук
	subRouter := a.router.NewRoute().Subrouter()
	subRouter.Use(checkcookiem.Middleware)

	subRouter.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUserInfo).Methods("GET")
	//subRouter.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	subRouter.HandleFunc("/auth", userHandler.GetLogin).Methods("GET")
	subRouter.HandleFunc("/users/{id:[0-9]+}", userHandler.DeleteUser).Methods("DELETE")
	subRouter.HandleFunc("/users/{id:[0-9]+}", userHandler.ChangeUserInfo).Methods("PATCH")

	// валидация всех данных, без кук
	a.router.HandleFunc("/users", userHandler.SignUp).Methods("POST")

	// валидация пароля
	a.router.HandleFunc("/login", userHandler.SignIn).Methods("POST")

	// не требуется
	a.router.HandleFunc("/login", userHandler.LogOut).Methods("DELETE")

	//a.router.HandleFunc("/users/{id:[0-9]+}/photos", userHandler.UploadPhoto).Methods("POST")
	//a.router.HandleFunc("/users/{id:[0-9]+}/photos", userHandler.DeletePhoto).Methods("DELETE")
}

func main() {
	a := App{}
	config := NewConfig()
	a.InitializeRoutes(config)
	err := a.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

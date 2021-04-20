package main

import (
	"math/rand"
	"net/http"
	"os"
	"server/internal/pickleapp/middleware"
	mainrep "server/internal/pickleapp/repository"
	"server/internal/pkg/chat/delivery"
	chatrep "server/internal/pkg/chat/repository"
	usecase2 "server/internal/pkg/chat/usecase"
	delivery2 "server/internal/pkg/message/delivery"
	mesrep "server/internal/pkg/message/repository"
	usecase3 "server/internal/pkg/message/usecase"
	"server/internal/pkg/models"
	delivery3 "server/internal/pkg/photo/delivery"
	usecase4 "server/internal/pkg/photo/usecase"
	"server/internal/pkg/session"
	sesrep "server/internal/pkg/session/repository"
	handler "server/internal/pkg/user/delivery"
	userrep "server/internal/pkg/user/repository"
	"server/internal/pkg/user/usecase"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"github.com/microcosm-cc/bluemonday"
	cors2 "github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type App struct {
	addr   string
	router *mux.Router
	Db     *sqlx.DB
	Logger *logrus.Logger
}

func (a *App) Start() error {
	a.Logger.Info("Server Start")

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

	//init config
	a.addr = currConfig.addr
	a.router = currConfig.router

	// init logger
	contextLogger := logrus.New()
	logrus.SetFormatter(&logrus.TextFormatter{})
	a.Logger = contextLogger

	// init db
	a.Db = mainrep.Init()
	userRep := userrep.UserRepository{DB: a.Db}
	sessionRep := sesrep.SessionRepository{DB: a.Db}
	messageRep := mesrep.MessageRepository{DB: a.Db}
	chatRep := chatrep.ChatRepository{DB: a.Db}

	clients := make(map[int]*websocket.Conn)
	// init uCases & handlers

	sanitizer := bluemonday.UGCPolicy()

	userUcase := usecase.UserUsecase{Db: userRep, Clients: &clients, Sanitizer: sanitizer}
	chatUcase := usecase2.ChatUsecase{Db: &chatRep, Clients: &clients}
	messUcase := usecase3.MessageUsecase{Db: &messageRep, Clients: &clients, Sanitizer: sanitizer}
	sessionManager := session.SessionsManager{DB: &sessionRep}

	chatHandler := delivery.ChatHandler{
		Db:       &chatRep,
		Sessions: &sessionManager,
		Usecase:  &chatUcase,
	}

	messHandler := delivery2.MessageHandler{
		Sessions: &sessionManager,
		Usecase:  &messUcase,
	}

	userHandler := handler.UserHandler{
		UserCase: &userUcase,
		Sessions: &sessionManager,
	}

	photousecase := usecase4.PhotoUseCase{
		Db: &userRep,
	}

	// init middlewares
	loggerm := middleware.LoggerMiddleware{
		Logger:  &models.Logger{Logger: logrus.NewEntry(a.Logger)},
		User:    &userUcase,
		Photo:   &photousecase,
		Session: &sessionManager,
	}

	photohandler := delivery3.PhotoHandler{Sessions: &sessionManager, Usecase: photousecase}

	checkcookiem := middleware.ValidateCookieMiddleware{Session: &sessionManager}

	a.router.Use(loggerm.Middleware)

	// validate cookie router
	subRouter := a.router.NewRoute().Subrouter()
	subRouter.Use(checkcookiem.Middleware)

	userHandler.SetHandlersWithCheckCookie(subRouter)
	userHandler.SetHandlersWithoutCheckCookie(a.router)
	photohandler.SetPhotoHandlers(subRouter)
	messHandler.SetMessageHandlers(subRouter)
	chatHandler.SetChatHandlers(subRouter)
}

func main() {
	a := App{}
	config := NewConfig()
	a.InitializeRoutes(config)
	err := a.Start()
	if err != nil {
		a.Logger.Error(err)
		os.Exit(1)
	}
}

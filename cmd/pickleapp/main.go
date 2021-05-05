package main

import (
	"math/rand"
	"net/http"
	"os"
	"server/internal/pickleapp/middleware"
	mainrep "server/internal/pickleapp/repository"
	"server/internal/pkg/chat/delivery"
	chatRepository "server/internal/pkg/chat/repository"
	chatUsecase "server/internal/pkg/chat/usecase"
	imageDelivery "server/internal/pkg/image/delivery"
	imageRepository "server/internal/pkg/image/repository"
	imageUsecase "server/internal/pkg/image/usecase"
	messageDelivery "server/internal/pkg/message/delivery"
	messageRepository "server/internal/pkg/message/repository"
	messageUsecase "server/internal/pkg/message/usecase"
	"server/internal/pkg/models"
	"server/internal/pkg/session"
	sessionRepository "server/internal/pkg/session/repository"
	userDelivery "server/internal/pkg/user/delivery"
	userRepository "server/internal/pkg/user/repository"
	"server/internal/pkg/user/usecase"
	"time"

	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"github.com/microcosm-cc/bluemonday"
	cors "github.com/rs/cors"
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

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://lepick.herokuapp.com", ""},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Access-Control-Allow-Headers", "Access-Control-Expose-Headers", "Authorization", "X-Requested-With", "X-Csrf-Token"},
		Debug:            false,
	})

	corsHandler := corsMiddleware.Handler(a.router)

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
	rand.Seed(time.Now().UnixNano())
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
	userRep := userRepository.UserRepository{DB: a.Db}
	sessionRep := sessionRepository.SessionRepository{DB: a.Db}
	messageRep := messageRepository.MessageRepository{DB: a.Db}
	chatRep := chatRepository.ChatRepository{DB: a.Db}

	imageRep := imageRepository.PostgresRepository{Db: a.Db}
	sess := awsSession.Must(awsSession.NewSession())
	awsRep := imageRepository.AwsImageRepository{
		Bucket:   "lepick-images",
		Svc:      s3.New(sess),
		Uploader: s3manager.NewUploader(sess),
	}

	clients := make(map[int]*websocket.Conn)
	// init uCases & handlers

	sanitizer := bluemonday.UGCPolicy()

	userUcase := usecase.UserUsecase{Db: &userRep, Clients: &clients, Sanitizer: sanitizer}
	chatUcase := chatUsecase.ChatUsecase{Db: &chatRep, Clients: &clients}
	messUcase := messageUsecase.MessageUsecase{Db: &messageRep, Clients: &clients, Sanitizer: sanitizer}
	sessionManager := session.SessionsManager{DB: &sessionRep}
	imageUcase := imageUsecase.ImageUsecase{
		Db:           &imageRep,
		ImageStorage: &awsRep,
	}

	chatHandler := delivery.ChatHandler{
		Sessions: &sessionManager,
		Usecase:  &chatUcase,
	}

	messHandler := messageDelivery.MessageHandler{
		Sessions: &sessionManager,
		Usecase:  &messUcase,
	}

	userHandler := userDelivery.UserHandler{
		UserCase: &userUcase,
		Sessions: &sessionManager,
	}

	imageHandler := imageDelivery.ImageHandler{
		Usecase:  &imageUcase,
		Sessions: &sessionManager,
	}

	// init middlewares
	loggerm := middleware.LoggerMiddleware{
		Logger:  &models.Logger{Logger: logrus.NewEntry(a.Logger)},
		User:    &userUcase,
		Image:   &imageUcase,
		Session: &sessionManager,
		Chat:    &chatUcase,
		Message: &messUcase,
	}

	checkcookiem := middleware.ValidateCookieMiddleware{Session: &sessionManager}

	rawRouter := a.router.NewRoute().Subrouter()

	userHandler.SetRawRouter(rawRouter)

	csrfRouter := a.router.NewRoute().Subrouter()

	csrfRouter.Use(loggerm.Middleware)
	csrfRouter.Use(middleware.CSRFMiddleware)

	// validate cookie router
	subRouter := csrfRouter.NewRoute().Subrouter()
	subRouter.Use(checkcookiem.Middleware)

	userHandler.SetHandlersWithCheckCookie(subRouter)
	userHandler.SetHandlersWithoutCheckCookie(csrfRouter)
	messHandler.SetMessageHandlers(subRouter)
	chatHandler.SetChatHandlers(subRouter)
	imageHandler.SetHandlers(subRouter)
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

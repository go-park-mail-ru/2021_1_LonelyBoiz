package main

import (
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"math/rand"
	"net/http"
	"os"
	"server/internal/pickleapp/middleware"
	mainrep "server/internal/pickleapp/repository"
	"server/internal/pkg/chat/delivery"
	chatrep "server/internal/pkg/chat/repository"
	usecase2 "server/internal/pkg/chat/usecase"
	imageDelivery "server/internal/pkg/image/delivery"
	imageRepository "server/internal/pkg/image/repository"
	imageUsecase "server/internal/pkg/image/usecase"
	delivery2 "server/internal/pkg/message/delivery"
	mesrep "server/internal/pkg/message/repository"
	usecase3 "server/internal/pkg/message/usecase"
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
		AllowedHeaders:   []string{"Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-Requested-With", "X-CSRF-Token"},
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

	userUcase := usecase.UserUsecase{Db: userRep, Clients: &clients, Sanitizer: sanitizer}
	chatUcase := usecase2.ChatUsecase{Db: chatRep, Clients: &clients}
	messUcase := usecase3.MessageUsecase{Db: messageRep, Clients: &clients, Sanitizer: sanitizer}
	imageUcase := imageUsecase.ImageUsecase{
		Db:           &imageRep,
		ImageStorage: &awsRep,
	}
	sessionManager := session.SessionsManager{DB: sessionRep}

	chatHandler := delivery.ChatHandler{
		Db:       chatRep,
		Sessions: &sessionManager,
		Usecase:  &chatUcase,
	}

	messHandler := delivery2.MessageHandler{
		Db:       messageRep,
		Sessions: &sessionManager,
		Usecase:  &messUcase,
	}

	userHandler := handler.UserHandler{
		UserCase: &userUcase,
		Sessions: &sessionManager,
	}

	imageHandler := imageDelivery.ImageHandler{
		Usecase:  &imageUcase,
		Sessions: &sessionManager,
	}

	// init middlewares
	loggerm := middleware.LoggerMiddleware{
		Logger:  contextLogger,
		User:    &userHandler,
		Chat:    &chatHandler,
		Message: &messHandler,
	}

	checkcookiem := middleware.ValidateCookieMiddleware{Session: &sessionManager}

	a.router.Use(loggerm.Middleware)

	// validate cookie router
	subRouter := a.router.NewRoute().Subrouter()
	subRouter.Use(checkcookiem.Middleware)

	// получить информацию о пользователе
	subRouter.HandleFunc("/users/{id:[0-9]+}", userHandler.GetUserInfo).Methods("GET")
	// получить ленту
	subRouter.HandleFunc("/feed", userHandler.GetUsers).Methods("GET")
	// првоерить куку
	subRouter.HandleFunc("/auth", userHandler.GetLogin).Methods("GET")
	// удалить юзера
	subRouter.HandleFunc("/users/{id:[0-9]+}", userHandler.DeleteUser).Methods("DELETE")
	// изменить информацию о юзере
	subRouter.HandleFunc("/users/{id:[0-9]+}", userHandler.ChangeUserInfo).Methods("PATCH")
	// поставить оценку юзеру из ленты
	subRouter.HandleFunc("/likes", userHandler.LikesHandler).Methods("POST")

	// загрузить новую фотку на сервер
	subRouter.HandleFunc("/images", imageHandler.UploadImage).Methods("POST")
	// удалить фотку
	subRouter.HandleFunc("/images/{uuid}", imageHandler.DeleteImage).Methods("DELETE")

	// валидация всех данных, без кук
	// регистрация
	a.router.HandleFunc("/users", userHandler.SignUp).Methods("POST")
	// логин
	a.router.HandleFunc("/login", userHandler.SignIn).Methods("POST")
	// логаут
	a.router.HandleFunc("/login", userHandler.LogOut).Methods("DELETE")

	// открытие вэсокетного соединения
	subRouter.HandleFunc("/ws", userHandler.WsHandler).Methods("GET")

	// получить чаты юзера
	subRouter.HandleFunc("/users/{userId:[0-9]+}/chats", chatHandler.GetChats).Methods("GET")

	// получить сообщения из чата
	subRouter.HandleFunc("/chats/{chatId:[0-9]+}/messages", messHandler.GetMessages).Methods("GET")
	// отправка нового сообщения
	subRouter.HandleFunc("/chats/{chatId:[0-9]+}/messages", messHandler.SendMessage).Methods("POST")
	// реакция
	subRouter.HandleFunc("/messages/{messageId:[0-9]+}", messHandler.ChangeMessage).Methods("PATCH")
	// отправка сообщения по вэбсокету собеседнику
	go messHandler.WebSocketMessageResponse()

	// отправка оповещения о новом чате
	go userHandler.WebSocketChatResponse()
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

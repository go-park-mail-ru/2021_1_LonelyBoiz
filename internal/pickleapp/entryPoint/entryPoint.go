package entryPoint

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"github.com/microcosm-cc/bluemonday"
	cors2 "github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"math/rand"
	"net/http"
	"os"
	session_proto2 "server/internal/auth_server/delivery/session"
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
	auth_server "server/internal/pkg/session"
	delivery4 "server/internal/pkg/session/delivery"
	"server/internal/pkg/session/repository"
	handler "server/internal/pkg/user/delivery"
	userrep "server/internal/pkg/user/repository"
	"server/internal/pkg/user/usecase"
	user_proto "server/internal/user_server/delivery/proto"
	"time"
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

func (a *App) InitializeRoutes(currConfig Config) []*grpc.ClientConn {

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
	sesRep := repository.SessionRepository{DB: a.Db}
	messageRep := mesrep.MessageRepository{DB: a.Db}
	chatRep := chatrep.ChatRepository{DB: a.Db}
	clients := make(map[int]*websocket.Conn)
	// init uCases & handlers
	sanitizer := bluemonday.UGCPolicy()

	//GRPC auth
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	authConn, err := grpc.Dial("localhost:5400", opts...)

	if err != nil {
		log.Print(1)
		grpclog.Fatalf("fail to dial: %v", err)
		panic(err)
	}

	client := session_proto2.NewAuthCheckerClient(authConn)

	//GRPC user
	opts = []grpc.DialOption{
		grpc.WithInsecure(),
	}

	userConn, err := grpc.Dial("localhost:5500", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
		panic(err)
	}

	userClient := user_proto.NewUserServiceClient(userConn)

	userUcase := usecase.UserUsecase{Db: &userRep, Clients: &clients, Sanitizer: sanitizer}
	chatUcase := usecase2.ChatUsecase{Db: &chatRep, Clients: &clients}
	messUcase := usecase3.MessageUsecase{Db: &messageRep, Clients: &clients, Sanitizer: sanitizer}

	chatHandler := delivery.ChatHandler{
		//Sessions: &sessionManager,
		Usecase: &chatUcase,
	}

	messHandler := delivery2.MessageHandler{
		//Sessions: &sessionManager,
		Usecase: &messUcase,
	}

	userHandler := handler.UserHandler{
		Server:   userClient,
		UserCase: &userUcase,
		Sessions: client,
	}

	photousecase := usecase4.PhotoUseCase{
		Db: &userRep,
	}

	authHandler := delivery4.AuthHandler{Usecase: &auth_server.SessionsManager{DB: &sesRep}}

	// init middlewares
	//csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth_server-key"))

	loggerm := middleware.LoggerMiddleware{
		Logger:  &models.Logger{Logger: logrus.NewEntry(a.Logger)},
		User:    &userUcase,
		Photo:   &photousecase,
		Chat:    &chatUcase,
		Message: &messUcase,
	}

	photohandler := delivery3.PhotoHandler{Usecase: photousecase}
	//
	checkcookiem := middleware.ValidateCookieMiddleware{Session: client}

	a.router.Use(loggerm.Middleware)
	a.router.Use(middleware.SetContextMiddleware)
	//a.router.Use(csrfMiddleware)
	//a.router.Use(middleware.CSRFMiddleware)

	// validate cookie router
	subRouter := a.router.NewRoute().Subrouter()
	subRouter.Use(checkcookiem.Middleware)

	a.router.HandleFunc("/login", authHandler.LogOut).Methods("DELETE")
	userHandler.SetHandlersWithCheckCookie(subRouter)
	userHandler.SetHandlersWithoutCheckCookie(a.router)
	photohandler.SetPhotoHandlers(subRouter)
	messHandler.SetMessageHandlers(subRouter)
	chatHandler.SetChatHandlers(subRouter)

	return []*grpc.ClientConn{userConn, authConn}
}

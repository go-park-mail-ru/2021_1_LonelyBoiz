package entryPoint

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	session_proto2 "server/internal/auth_server/delivery/session"
	"server/internal/pickleapp/middleware"
	"server/internal/pickleapp/repository"
	"server/internal/pkg/chat/delivery"
	chatRepository "server/internal/pkg/chat/repository"
	chatUsecase "server/internal/pkg/chat/usecase"
	messageDelivery "server/internal/pkg/message/delivery"
	messageRepository "server/internal/pkg/message/repository"
	messageUsecase "server/internal/pkg/message/usecase"
	"server/internal/pkg/models"
	"server/internal/pkg/session"
	sessionRepository "server/internal/pkg/session/repository"
	userDelivery "server/internal/pkg/user/delivery"
	userRepository "server/internal/pkg/user/repository"
	"server/internal/pkg/user/usecase"
	user_proto "server/internal/user_server/delivery/proto"
	"time"

	authHandler "server/internal/pkg/session/delivery"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"github.com/microcosm-cc/bluemonday"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	cors2 "github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type App struct {
	Addr   string
	Router *mux.Router
	Db     *sqlx.DB
	Logger *logrus.Logger
}

func (a *App) Start() error {
	a.Logger.Info("Server Start")

	prometheus.MustRegister(middleware.FooCount, middleware.Hits)

	cors := cors2.New(cors2.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://lepick.online"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Access-Control-Allow-Headers", "Access-Control-Expose-Headers", "Access-Control-Allow-Origin", "Authorization", "X-Requested-With", "X-CSRF-Token", "Server"},
		Debug:            false,
	})

	corsHandler := cors.Handler(a.Router)

	s := http.Server{
		Addr:         a.Addr,
		Handler:      corsHandler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	//err := s.ListenAndServe()
	err := s.ListenAndServeTLS(os.Getenv("SSL_PUBLIC"), os.Getenv("SSL_PRIVATE"))
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
	a.Addr = currConfig.addr
	a.Router = currConfig.router

	// init logger
	contextLogger := logrus.New()
	logrus.SetFormatter(&logrus.TextFormatter{})
	a.Logger = contextLogger

	// init db
	a.Db = repository.Init()
	userRep := userRepository.UserRepository{DB: a.Db}
	sessionRep := sessionRepository.SessionRepository{DB: a.Db}
	messageRep := messageRepository.MessageRepository{DB: a.Db}
	chatRep := chatRepository.ChatRepository{DB: a.Db}

	clients := make(map[int]*websocket.Conn)

	//GRPC auth
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	authConn, err := grpc.Dial("auth:5400", opts...)

	if err != nil {
		log.Print(1)
		grpclog.Fatalf("fail to dial: %v", err)
		panic(err)
	}

	authClient := session_proto2.NewAuthCheckerClient(authConn)

	//GRPC user
	opts = []grpc.DialOption{
		grpc.WithInsecure(),
	}

	userConn, err := grpc.Dial("user:5500", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
		panic(err)
	}

	userClient := user_proto.NewUserServiceClient(userConn)

	// init uCases & handlers
	sanitizer := bluemonday.UGCPolicy()
	userUcase := usecase.UserUsecase{Db: &userRep, Clients: &clients, Sanitizer: sanitizer}
	chatUcase := chatUsecase.ChatUsecase{Db: &chatRep}
	messUcase := messageUsecase.MessageUsecase{Db: &messageRep, Clients: &clients, Sanitizer: sanitizer}
	sessionManager := session.SessionsManager{DB: &sessionRep}

	chatHandler := delivery.ChatHandler{
		Usecase: &chatUcase,
		Server:  userClient,
	}

	messHandler := messageDelivery.MessageHandler{
		Usecase: &messUcase,
		Server:  userClient,
	}

	userHandler := userDelivery.UserHandler{
		Server:   userClient,
		UserCase: &userUcase,
		Sessions: authClient,
	}

	authHandler := authHandler.AuthHandler{
		Usecase: &sessionManager,
	}

	// init middlewares
	loggerm := middleware.LoggerMiddleware{
		Logger: &models.Logger{Logger: logrus.NewEntry(a.Logger)},
		User:   &userUcase,
		//Session: &sessionManager,
		Chat:    &chatUcase,
		Message: &messUcase,
	}

	checkcookiem := middleware.ValidateCookieMiddleware{
		Session: authClient,
	}

	a.Router.Handle("/metrics", promhttp.Handler()).Methods("GET")

	rawRouter := a.Router.NewRoute().Subrouter()

	userHandler.SetRawRouter(rawRouter)

	csrfRouter := a.Router.NewRoute().Subrouter()

	csrfRouter.Use(loggerm.Middleware)
	//csrfRouter.Use(middleware.CSRFMiddleware)
	csrfRouter.Use(middleware.SetContextMiddleware)

	// validate cookie router
	subRouter := csrfRouter.NewRoute().Subrouter()
	subRouter.Use(checkcookiem.Middleware)

	userHandler.SetHandlersWithCheckCookie(subRouter)
	userHandler.SetHandlersWithoutCheckCookie(csrfRouter)
	messHandler.SetMessageHandlers(subRouter)
	chatHandler.SetChatHandlers(subRouter)
	authHandler.SetAuthHandler(csrfRouter)

	return []*grpc.ClientConn{userConn, authConn}
}

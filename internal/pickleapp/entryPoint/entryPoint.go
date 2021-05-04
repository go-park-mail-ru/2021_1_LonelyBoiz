package entryPoint

import (
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
	"server/internal/pickleapp/repository"
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
	a.Db = repository.Init()
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

	authClient := session_proto2.NewAuthCheckerClient(authConn)

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
		Usecase: &chatUcase,
	}

	messHandler := messageDelivery.MessageHandler{
		Usecase: &messUcase,
	}

	userHandler := userDelivery.UserHandler{
		Server:   userClient,
		UserCase: &userUcase,
		Sessions: authClient,
	}

	imageHandler := imageDelivery.ImageHandler{
		Usecase:  &imageUcase,
		Sessions: &sessionManager,
	}

	// init middlewares
	loggerm := middleware.LoggerMiddleware{
		Logger: &models.Logger{Logger: logrus.NewEntry(a.Logger)},
		User:   &userUcase,
		Image:  &imageUcase,
		//Session: &sessionManager,
		Chat:    &chatUcase,
		Message: &messUcase,
	}

	checkcookiem := middleware.ValidateCookieMiddleware{
		Session: authClient,
	}

	a.router.Use(loggerm.Middleware)
	//a.router.Use(middleware.CSRFMiddleware)
	a.router.Use(middleware.SetContextMiddleware)

	// validate cookie router
	subRouter := a.router.NewRoute().Subrouter()
	subRouter.Use(checkcookiem.Middleware)

	userHandler.SetHandlersWithCheckCookie(subRouter)
	userHandler.SetHandlersWithoutCheckCookie(a.router)
	messHandler.SetMessageHandlers(subRouter)
	chatHandler.SetChatHandlers(subRouter)
	imageHandler.SetHandlers(subRouter)

	return []*grpc.ClientConn{userConn, authConn}
}

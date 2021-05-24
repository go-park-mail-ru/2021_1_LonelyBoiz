package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	session_proto2 "server/internal/auth_server/delivery/session"
	"server/internal/pickleapp/entryPoint"
	"server/internal/pickleapp/middleware"
	"server/internal/pickleapp/repository"
	imageDelivery "server/internal/pkg/image/delivery"
	imageRepository "server/internal/pkg/image/repository"
	imageUsecase "server/internal/pkg/image/usecase"
	"server/internal/pkg/models"
	"server/internal/pkg/utils/metrics"
	"time"

	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gocv.io/x/gocv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Config struct {
	addr    string
	userIds int
	router  *mux.Router
}

func NewConfig() Config {
	rand.Seed(time.Now().UnixNano())
	newConfig := Config{}
	port := "7000"
	newConfig.addr = ":" + port
	newConfig.userIds = 0
	newConfig.router = mux.NewRouter()
	return newConfig
}

func main() {
	a := entryPoint.App{}
	config := NewConfig()

	//init config
	a.Addr = config.addr
	a.Router = config.router

	// init logger
	contextLogger := logrus.New()
	logrus.SetFormatter(&logrus.TextFormatter{})
	a.Logger = contextLogger

	// init db
	a.Db = repository.Init()
	sess := awsSession.Must(awsSession.NewSession())
	awsRep := imageRepository.AwsImageRepository{
		Bucket:   "lepick-images",
		Svc:      s3.New(sess),
		Uploader: s3manager.NewUploader(sess),
	}

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
	defer authConn.Close()

	authClient := session_proto2.NewAuthCheckerClient(authConn)

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load("/Users/nick_nak/programs/2_parkMail/2021_1_LonelyBoiz/cmd/image_server/haarcascade_frontalface_default.xml") {
		fmt.Printf("Error reading cascade file:")
		return
	}

	// init uCases & handlers
	imageUcase := imageUsecase.ImageUsecase{
		Db:            &imageRepository.PostgresRepository{Db: a.Db},
		ImageStorage:  &awsRep,
		FaceDetection: classifier,
	}

	imageHandler := imageDelivery.ImageHandler{
		Usecase: &imageUcase,
	}
	// init middlewares
	loggerm := middleware.LoggerMiddleware{
		Logger: &models.Logger{Logger: logrus.NewEntry(a.Logger)},
		Image:  &imageUcase,
	}

	checkcookiem := middleware.ValidateCookieMiddleware{
		Session: authClient,
	}

	metrics.New()

	csrfRouter := a.Router.NewRoute().Subrouter()

	csrfRouter.Use(loggerm.Middleware)
	csrfRouter.Use(middleware.SetContextMiddleware)

	// validate cookie router
	subRouter := csrfRouter.NewRoute().Subrouter()
	subRouter.Use(checkcookiem.Middleware)

	imageHandler.SetHandlers(subRouter)

	err = a.Start()
	if err != nil {
		a.Logger.Error(err)
		os.Exit(1)
	}
}

package main

import (
	"context"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"log"
	"math/rand"
	"net"
	delivery2 "server/internal/image_server/delivery"
	imageProto "server/internal/image_server/delivery/proto"
	"server/internal/pickleapp/repository"
	imageRepository "server/internal/pkg/image/repository"
	"server/internal/pkg/image/usecase"
	"server/internal/pkg/models"
	"time"
)

type ServerInterceptor struct {
	Logger *models.Logger
}

func (s *ServerInterceptor) logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()

	md, _ := metadata.FromIncomingContext(ctx)

	reqIds := md.Get("requestId")
	var reqId string
	if len(reqIds) != 0 {
		reqId = reqIds[0]
	}

	s.Logger.Logger = s.Logger.Logger.WithFields(logrus.Fields{
		"server":    "[AUTH]",
		"requestId": reqId,
		"method":    info.FullMethod,
		"context":   md,
		"request":   req,
		"response":  resp,
		"error":     err,
		"work_time": time.Since(start),
	})

	reply, err := handler(ctx, req)

	s.Logger.LogInfo("Auth Interceptor")
	return reply, err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	listener, err := net.Listen("tcp", ":5200")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	ServerInterceptor := ServerInterceptor{&models.Logger{Logger: logrus.NewEntry(logrus.StandardLogger())}}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(ServerInterceptor.logger))

	sess := awsSession.Must(awsSession.NewSession())

	awsRep := imageRepository.AwsImageRepository{
		Bucket:   "lepick-images",
		Svc:      s3.New(sess),
		Uploader: s3manager.NewUploader(sess),
	}

	imageServer := delivery2.ImageServer{
		Usecase: &usecase.ImageUsecase{
			Db:              &imageRepository.PostgresRepository{Db: repository.Init()},
			ImageStorage:    &awsRep,
			LoggerInterface: ServerInterceptor.Logger,
		},
	}

	imageProto.RegisterImageServiceServer(grpcServer, &imageServer)
	log.Print("Image Server START at 5200")
	err = grpcServer.Serve(listener)

	if err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
		return
	}
}

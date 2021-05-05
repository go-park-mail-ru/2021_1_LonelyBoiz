package main

import (
	"github.com/gorilla/websocket"
	"github.com/microcosm-cc/bluemonday"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"log"
	"math/rand"
	"net"
	session_proto2 "server/internal/auth_server/delivery/session"
	"server/internal/pickleapp/repository"
	chatRepository "server/internal/pkg/chat/repository"
	chatUsecase "server/internal/pkg/chat/usecase"
	repository3 "server/internal/pkg/message/repository"
	messageUsecase "server/internal/pkg/message/usecase"
	"server/internal/pkg/models"
	repository2 "server/internal/pkg/user/repository"
	"server/internal/pkg/user/usecase"
	delivery2 "server/internal/user_server/delivery"
	userProto "server/internal/user_server/delivery/proto"
	"time"
)

type UserServerInterceptor struct {
	Logger *models.Logger
}

func (s *UserServerInterceptor) logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)

	reqIds := md.Get("requestId")
	var reqId string
	if len(reqIds) != 0 {
		reqId = reqIds[0]
	}

	s.Logger.Logger = s.Logger.Logger.WithFields(logrus.Fields{
		"server":    "[USER]",
		"requestId": reqId,
		"method":    info.FullMethod,
		"context":   md,
		"request":   req,
		"response":  resp,
		"error":     err,
		"work_time": time.Since(start),
	})

	reply, err := handler(ctx, req)

	s.Logger.LogInfo("USER Interceptor")
	return reply, err
}

func main() {
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

	auth := session_proto2.NewAuthCheckerClient(authConn)

	// main part
	rand.Seed(time.Now().UnixNano())
	listener, err := net.Listen("tcp", ":5500")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	ServerInterceptor := UserServerInterceptor{&models.Logger{Logger: logrus.NewEntry(logrus.StandardLogger())}}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(ServerInterceptor.logger))

	db := repository.Init()

	clients := make(map[int]*websocket.Conn)

	userServer := delivery2.UserServer{
		UserUsecase: &usecase.UserUsecase{
			Db:              &repository2.UserRepository{DB: db},
			LoggerInterface: ServerInterceptor.Logger,
			Sanitizer:       bluemonday.NewPolicy(),
			Clients:         &clients,
		},
		ChatUsecase: &chatUsecase.ChatUsecase{
			Clients:         &clients,
			LoggerInterface: ServerInterceptor.Logger,
			Db:              &chatRepository.ChatRepository{DB: db},
		},
		MessageUsecase: &messageUsecase.MessageUsecase{
			Clients:         &clients,
			Db:              &repository3.MessageRepository{DB: db},
			LoggerInterface: ServerInterceptor.Logger,
			Sanitizer:       bluemonday.NewPolicy(),
		},
		Sessions: auth,
	}

	userProto.RegisterUserServiceServer(grpcServer, &userServer)
	log.Print("User Server START at 5500")
	err = grpcServer.Serve(listener)

	if err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
		return
	}
}

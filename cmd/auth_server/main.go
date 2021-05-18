package main

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"log"
	"math/rand"
	"net"
	"server/internal/auth_server/delivery"
	"server/internal/auth_server/delivery/session"
	"server/internal/pickleapp/repository"
	"server/internal/pkg/models"
	authserver "server/internal/pkg/session"
	sesrep "server/internal/pkg/session/repository"
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
	listener, err := net.Listen("tcp", ":5400")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	ServerInterceptor := ServerInterceptor{&models.Logger{Logger: logrus.NewEntry(logrus.StandardLogger())}}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(ServerInterceptor.logger))

	authServer := delivery.AuthServer{
		Usecase: &authserver.SessionsManager{
			Logger: ServerInterceptor.Logger,
			DB:     &sesrep.SessionRepository{DB: repository.Init()},
		},
	}

	session_proto.RegisterAuthCheckerServer(grpcServer, &authServer)
	log.Print("Auth Server START at 5400")
	err = grpcServer.Serve(listener)

	if err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
		return
	}
}

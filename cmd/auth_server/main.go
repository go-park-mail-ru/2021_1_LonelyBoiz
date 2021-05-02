package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"math/rand"
	"net"
	"server/internal/auth_server/delivery"
	"server/internal/auth_server/delivery/session"
	"server/internal/pickleapp/repository"
	authserver "server/internal/pkg/session"
	sesrep "server/internal/pkg/session/repository"
	"time"
)

func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("hi from interceptor")
	log.Println(info.FullMethod)
	h, err := handler(ctx, req)

	return h, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	listener, err := net.Listen("tcp", ":5400")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	authServer := delivery.AuthServer{
		Usecase: &authserver.SessionsManager{
			DB: &sesrep.SessionRepository{DB: repository.Init()},
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

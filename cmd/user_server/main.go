package main

import (
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
	"server/internal/pkg/models"
	repository2 "server/internal/pkg/user/repository"
	"server/internal/pkg/user/usecase"
	delivery2 "server/internal/user_server/delivery"
	userProto "server/internal/user_server/delivery/proto"
	"time"
)

func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log.Println("hi from interceptor")

	d, _ := metadata.FromIncomingContext(ctx)

	log.Println(d)
	h, err := handler(ctx, req)

	return h, nil
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

	grpcServer := grpc.NewServer()

	db := repository.Init()

	loger := logrus.Logger{}

	userServer := delivery2.UserServer{
		Usecase: &usecase.UserUsecase{
			Db:              &repository2.UserRepository{DB: db},
			LoggerInterface: &models.Logger{Logger: loger.WithField("key", "value")},
			Sanitizer:       bluemonday.NewPolicy(),
		},
		Sessions: auth,
	}

	userProto.RegisterUserServiceServer(grpcServer, &userServer)
	log.Print("Auth Server START at 5500")
	err = grpcServer.Serve(listener)

	if err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
		return
	}
}

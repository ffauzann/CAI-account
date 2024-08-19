package grpc

import (
	"github.com/ffauzann/CAI-account/internal/service"
	"github.com/ffauzann/CAI-account/proto/gen"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type srv struct {
	gen.UnimplementedAuthServiceServer
	gen.UnimplementedUserServiceServer
	gen.UnimplementedAccountServiceServer
	service service.Service
	logger  *zap.Logger
}

func New(server *grpc.Server, userSrv service.Service, logger *zap.Logger) {
	srv := srv{
		service: userSrv,
		logger:  logger,
	}
	gen.RegisterAuthServiceServer(server, &srv)
	gen.RegisterUserServiceServer(server, &srv)
	gen.RegisterAccountServiceServer(server, &srv)
	reflection.Register(server)
}

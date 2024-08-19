package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	authInterceptor "github.com/ffauzann/common/interceptor/grpc/unary/authentication"
	ctxTagsInterceptor "github.com/ffauzann/common/interceptor/grpc/unary/ctxtags"
	logInterceptor "github.com/ffauzann/common/interceptor/grpc/unary/logging"
	recoveryInterceptor "github.com/ffauzann/common/interceptor/grpc/unary/recovery"

	grpcCtxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/ffauzann/CAI-account/client"
	"github.com/ffauzann/CAI-account/internal/constant"
	deliveryGRPC "github.com/ffauzann/CAI-account/internal/delivery/grpc"
	deliveryHTTP "github.com/ffauzann/CAI-account/internal/delivery/http"
	"github.com/ffauzann/CAI-account/internal/repository"
	"github.com/ffauzann/CAI-account/internal/service"
	"github.com/ffauzann/CAI-account/internal/util"
	"github.com/ffauzann/CAI-account/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

type Server struct {
	GRPC   GRPC
	HTTP   HTTP
	Logger Logger
}

type GRPC struct {
	Address string
	Port    uint32
	Server  *grpc.Server
}

type HTTP struct {
	Address string
	Port    uint32
	Timeout string
	Server  *http.Server
}

func (c *Config) StartServer() {
	var wg sync.WaitGroup
	wg.Add(2) //nolint

	// Init repo
	dbRepo := repository.NewDB(c.Database.SQL.DB, c.App, c.Server.Logger.Zap)
	redisRepo := repository.NewRedis(c.Cache.Redis.Client, c.App, c.Server.Logger.Zap)
	whatsappClientRepo := repository.NewWhatsappClient(c.App, c.Server.Logger.Zap)

	// Init service
	svc := service.New(dbRepo, redisRepo, whatsappClientRepo, c.App, c.Server.Logger.Zap)

	go func() {
		defer wg.Done()
		c.startGRPCServer(svc)
	}()

	go func() {
		defer wg.Done()
		c.startHTTPProxyServer(svc)
	}()

	// Graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
	fmt.Println("Gracefully shutting down...")

	c.Server.GRPC.Server.GracefulStop()
	fmt.Println("gRPC server has been shutdown.")
	c.Server.HTTP.Server.Shutdown(context.Background())
	fmt.Println("HTTP proxy server has been shutdown.")
}

func (c *Config) startGRPCServer(svc service.Service) {
	addr := fmt.Sprintf("%s:%d", c.Server.GRPC.Address, c.Server.GRPC.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
		return
	}

	interceptors := []grpc.UnaryServerInterceptor{
		ctxTagsInterceptor.UnaryServerInterceptor(),
		recoveryInterceptor.UnaryServerInterceptor(c.Server.Logger.Zap),
		// // authInterceptor can be removed once it is handled by API gateway.
		// authInterceptor.UnaryServerInterceptor(
		// 	&authInterceptor.Config{
		// 		Iss:        c.App.Jwt.AccessToken.Iss,
		// 		Alg:        jwt.SigningMethodHS256.Alg(),
		// 		SigningKey: c.App.Jwt.AccessToken.SigningKey,
		// 	},
		// 	authInterceptor.WithCustomMetadataKey("Authorization"),
		// 	authInterceptor.WithCustomClaims(&client.Claims{}),
		// 	authInterceptor.WithExcludedMethods(c.App.Auth.ExcludedMethods...),
		// ),
		authInterceptor.UnaryServerInterceptor(
			&authInterceptor.Config{
				Iss:     c.App.Jwt.AccessToken.Iss,
				Alg:     authInterceptor.AlgRS256,
				JwksURL: fmt.Sprintf("http://%s:%d%s", c.Server.HTTP.Address, c.Server.HTTP.Port, client.JwksPath),
			},
			authInterceptor.WithCustomMetadataKey("Authorization"),
			authInterceptor.WithCustomClaims(&client.Claims{}),
			authInterceptor.WithExcludedMethods(c.App.Auth.ExcludedMethods...),
		),
		logInterceptor.UnaryServerInterceptor(
			c.Server.Logger.Zap,
			logInterceptor.WithErrorParser(constant.MapGRPCErrCodes),
			logInterceptor.WithCtxTag(true),
		),
		grpcCtxtags.UnaryServerInterceptor(),
	}
	opts := grpc.ChainUnaryInterceptor(interceptors...)
	c.Server.GRPC.Server = grpc.NewServer(opts)

	deliveryGRPC.New(c.Server.GRPC.Server, svc, c.Server.Logger.Zap)
	fmt.Printf("gRPC server started on %s\n", addr)

	if err := c.Server.GRPC.Server.Serve(lis); err != nil {
		log.Fatal(err)
		return
	}
}

func (c *Config) startHTTPProxyServer(svc service.Service) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcMux := runtime.NewServeMux(
		runtime.WithMetadata(
			func(ctx context.Context, r *http.Request) metadata.MD {
				return metadata.Pairs("X-Forwarded-Method", r.Method)
			},
		),
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:     true,
					EmitDefaultValues: true,
				},
			},
		),
	)

	timeout, _ := time.ParseDuration(c.Server.HTTP.Timeout)
	addr := fmt.Sprintf("%s:%d", c.Server.HTTP.Address, c.Server.HTTP.Port)
	c.Server.HTTP.Server = &http.Server{
		Addr:              addr,
		Handler:           grpcMux,
		ReadHeaderTimeout: timeout,
		ReadTimeout:       timeout,
		WriteTimeout:      timeout,
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := gen.RegisterAuthServiceHandlerFromEndpoint(ctx, grpcMux, fmt.Sprintf("%s:%d", c.Server.GRPC.Address, c.Server.GRPC.Port), opts); err != nil {
		util.Log().Error(err.Error())
		return
	}

	if err := gen.RegisterUserServiceHandlerFromEndpoint(ctx, grpcMux, fmt.Sprintf("%s:%d", c.Server.GRPC.Address, c.Server.GRPC.Port), opts); err != nil {
		util.Log().Error(err.Error())
		return
	}

	if err := gen.RegisterAccountServiceHandlerFromEndpoint(ctx, grpcMux, fmt.Sprintf("%s:%d", c.Server.GRPC.Address, c.Server.GRPC.Port), opts); err != nil {
		util.Log().Error(err.Error())
		return
	}

	deliveryHTTP.New(grpcMux, svc)

	fmt.Printf("HTTP proxy server started on %s\n", addr)

	if err := c.Server.HTTP.Server.ListenAndServe(); err != nil {
		util.Log().Error(err.Error())
		return
	}
}

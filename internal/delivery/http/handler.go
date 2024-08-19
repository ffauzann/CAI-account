package http

import (
	"net/http"

	"github.com/ffauzann/CAI-account/internal/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type srv struct {
	service service.Service
}

func New(server *runtime.ServeMux, userSrv service.Service) {
	s := &srv{
		service: userSrv,
	}

	server.HandlePath(http.MethodGet, "/user/api/v1/r/utilities/healthz", s.Health)
	server.HandlePath(http.MethodGet, "/user/api/v1/r/utilities/.well-known/jwks.json", s.Jwks)
}

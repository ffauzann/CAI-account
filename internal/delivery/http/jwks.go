package http

import (
	"encoding/json"
	"net/http"

	"github.com/ffauzann/CAI-account/internal/util"
)

func (s *srv) Jwks(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	ctx := r.Context()
	jwks, err := s.service.Jwks(ctx)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	b, err := json.Marshal(jwks)
	if err != nil {
		util.LogContext(ctx).Error(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(b))
}

package http

import "net/http"

func (s *srv) Health(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"OK"}`))
}

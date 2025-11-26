package api

import (
	"net/http"
)

type Server struct{}

func (s *Server) RegisterRoutes() {
	mux.HandleFunc("/health", s.HealthCheckHandler)
}

// HealthCheckHandler handles the health check endpoint.
func (s *Server) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

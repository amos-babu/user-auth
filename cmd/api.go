package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddr string
}

func NewApiServer(listenAddr string) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/health", s.HealthCheckHandler).Methods("GET")

	log.Println("Server running at port:", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, router)
}

func (s *ApiServer) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

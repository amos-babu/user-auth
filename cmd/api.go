package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddr string
	config     Config
}

type Config struct {
	cfg string
}

func NewApiServer(listenAddr string) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
	}
}

func (s *ApiServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/health", s.HealthCheckHandler).Methods("GET")
	router.HandleFunc("/register", s.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", s.LoginHandler).Methods("POST")

	log.Println("Server running at port:", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, router)
}

func (s *ApiServer) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Status string `json:"status"`
	}{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (s *ApiServer) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Registration logic goes here
}

func (s *ApiServer) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Login logic goes here
}

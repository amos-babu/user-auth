package main

import (
	"github.com/amos-babu/user-auth/api"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	apiServer := &api.Server{}
	apiServer.RegisterRoutes()
}

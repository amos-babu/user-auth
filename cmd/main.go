package main

import "github.com/amos-babu/user-auth/internals/env"

func main() {
	Addr := env.GetString("ADDR", ":3000")
	server := NewApiServer(Addr)
	if err := server.Run(); err != nil {
		panic("Server failed to start: " + err.Error())
	}

}

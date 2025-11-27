package main

func main() {
	server := NewApiServer(":8080")
	if err := server.Run(); err != nil {
		panic("Server failed to start: " + err.Error())
	}

}

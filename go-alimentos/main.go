package main

import "go-alimentos/internal/app/server"

func main() {
	server := server.NewServer()
	server.Run()
}

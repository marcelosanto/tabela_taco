package main

import (
	"go-alimentos/server"
)

func main() {
	server := server.NewServer()
	server.Run()
}

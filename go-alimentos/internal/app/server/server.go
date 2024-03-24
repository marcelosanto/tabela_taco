package server

import (
	"go-alimentos/internal/app/server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	gin.SetMode(gin.ReleaseMode)

	return Server{
		port:   ":5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Print("Server is running on port", s.port)
	log.Fatal(router.Run(s.port))
}

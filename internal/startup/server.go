package startup

import (
	"context"
	"github.com/GGmaz/BookManager/config"
	v1 "github.com/GGmaz/BookManager/internal/api/v1"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	r := gin.Default()
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	v1.RegisterVersion(r, context.Background())

	err := r.Run(":" + server.config.Port)
	if err != nil {
		log.Fatal("Could not start server: " + err.Error())
		return
	}

	println("Starting server on port: " + server.config.Port)
}

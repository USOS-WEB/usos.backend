package server

import (
	"github.com/USOS-WEB/usos.backend/config"
	"github.com/USOS-WEB/usos.backend/database"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     config.Config
	db      database.Database
	router     *gin.Engine
}

func NewServer(config config.Config, db database.Database) (*Server, error) {
	server := &Server{
		config:     config,
		db:      db,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/route", server.getRoute)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

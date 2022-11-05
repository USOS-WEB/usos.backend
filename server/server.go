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

	router.Use(CORSMiddleware())

	router.GET("/", server.nothing)
	router.POST("/route", server.getRoute)
	router.GET("/points", server.getAllPoints)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

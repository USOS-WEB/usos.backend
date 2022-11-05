package main

import (
	"log"
	"os"

	"github.com/USOS-WEB/usos.backend/config"
	"github.com/USOS-WEB/usos.backend/database"
	"github.com/USOS-WEB/usos.backend/server"
	"github.com/go-pg/pg/v10"
)

func main() {
	config, err := config.LoadConfigFromFile(".env")
	if err != nil {
		log.Fatal("cannot load config")
	}

	db := database.Connect(pg.Options{
		Addr:     config.DB_HOST + ":" + config.DB_PORT,
		User:     config.DB_USER,
		Password: config.DB_PASSWORD,
		Database: config.DB_DBNAME,
	})

	if err := db.CheckConnection(); err != nil{
		log.Fatal("cannot connect to database")
	}

	runHTTPServer(config, db)
}


func runHTTPServer(config config.Config, db database.Database) {
	server, err := server.NewServer(config, db)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("cannot start server")
	}
}

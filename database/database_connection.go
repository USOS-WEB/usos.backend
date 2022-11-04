package database

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

func Connect(config pg.Options) Database {
	db := pg.Connect(&config)
	if err:= db.Ping(context.Background()); err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Connected to database")
	}

	return &data{db}
}

func (db *data) CheckConnection() error {
	if err:= db.DB.Ping(context.Background()); err != nil{
		return fmt.Errorf("No connection to database: %s", err)
	}

	return nil
}

func (db *data) Close() {
	db.DB.Close()
}

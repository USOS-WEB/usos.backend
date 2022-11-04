package database

import (
	"context"
	"fmt"
	"log"

	"github.com/go-pg/pg/v10"
)

func (db *data) Test() {

	fmt.Println("byłem tu")

	edge := &Edge{}

	point := &Point{}

	err := db.Model(edge).Select()
	if err != nil {
		log.Panic(err)
	}

	err = db.Model(point).Where("id = ?", edge.Point1).Select()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(point)

	err = db.Model(point).Where("id = ?", edge.Point2).Select()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(point)
}

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

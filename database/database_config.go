package database

import (
	"github.com/go-pg/pg/v10"
)

type data struct {
	*pg.DB
}

type Database interface {
	//General
	Close()
	CheckConnection() error

	Test()

	PointSelectByID(id string) (*Point, error)
	GetAllEdgePoints() (edgePoints []EdgePoints, err error)
}

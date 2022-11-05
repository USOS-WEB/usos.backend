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
	FloorIdSelectByPointId(id string) (*[]Point_floor, error)
	FloorSelect() (*[]Floor, error)
}

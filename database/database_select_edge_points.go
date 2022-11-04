package database

import (
	"log"
)

type EdgePoints struct {
	Point1 Point
	Point2 Point
}

func (db *data) GetAllEdgePoints() (edgePoints []EdgePoints, err error) {
	edges := &[]Edge{}

	err = db.Model(edges).Select()
	if err != nil {
		log.Panic(err)
	}

	for _, edge := range *edges {
		p1, err := db.PointSelectByID(edge.Point1)
		if err != nil {
			log.Panic(err)
		}

		p2, err := db.PointSelectByID(edge.Point2)
		if err != nil {
			log.Panic(err)
		}

		edgePoints = append(edgePoints, EdgePoints{
			Point1: *p1,
			Point2: *p2,
		})
	}

	return
}

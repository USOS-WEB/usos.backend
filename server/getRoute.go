package server

import (
	"fmt"
	"net/http"

	"github.com/USOS-WEB/usos.backend/database"
	"github.com/gin-gonic/gin"
	"gopkg.in/karalabe/cookiejar.v2/graph"
	"gopkg.in/karalabe/cookiejar.v2/graph/bfs"
)

func getPointIndex(array []database.Point, id string) (int){
	for i := range array {
		if array[i].Id == id {
			return i
		}
	}

	return -1
}

func (s *Server) getRoute(ctx *gin.Context) {
	var req getRouteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	edgePoints, err := s.db.GetAllEdgePoints()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	points := []database.Point{}

	for _, edge := range edgePoints {
		idx1 := getPointIndex(points, edge.Point1.Id)
		if idx1 == -1 {
			points = append(points, edge.Point1)
		}

		idx2 := getPointIndex(points, edge.Point2.Id)
		if idx2 == -1 {
			points = append(points, edge.Point2)
		}
	}

	g := graph.New(len(points))

	for _, edge := range edgePoints {
		g.Connect(getPointIndex(points, edge.Point1.Id), getPointIndex(points, edge.Point2.Id))
	}

	b := bfs.New(g, getPointIndex(points, req.Start))

	path := b.Path(getPointIndex(points, req.Stop))

	for _,i := range path {
		fmt.Println(points[i])
	}
}

type getRouteRequest struct {
	Start string `json:"start"`
	Stop string `json:"stop"`
}

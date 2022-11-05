package server

import (
	"net/http"

	"github.com/USOS-WEB/usos.backend/database"
	"github.com/gin-gonic/gin"
)

func (s *Server) getAllPoints(ctx *gin.Context) {
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

	res := getPointsResponse{
		Points: points,
	}

	ctx.JSON(http.StatusOK, res)
}

type getPointsResponse struct {
	Points []database.Point `json:"points"`
}

package server

import (
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

	responsePoints := []ResponsePoint{}

	for _, i := range path {
		floors, err := s.db.FloorIdSelectByPointId(points[i].Id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		floorsDone := []string{}

		for _,floor := range *floors{
			floorsDone = append(floorsDone, floor.Floor_id)
		}

		responsePoint := ResponsePoint{
			Id:          points[i].Id,
			Name:        points[i].Name,
			Description: points[i].Description,
			Floor_areas: points[i].Floor_areas,
			Url:         points[i].Url,
			Width:       points[i].Width,
			Height:      points[i].Height,
			Floors:      floorsDone,
		}
		responsePoints = append(responsePoints, responsePoint)
	}

	allFloors, err := s.db.FloorSelect()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	floorsToSend := []ResponseFloor{}

	for _,floor := range *allFloors{
		floorsToSend = append(floorsToSend, ResponseFloor{
			Id:    floor.Id,
			Name:  floor.Name,
			Image: Image{
				Url:    floor.Url,
				Width:  floor.Width,
				Height: floor.Height,
			},
		})
	}

	res := getRouteResponse{
		Path:   responsePoints,
		Floors: floorsToSend,
	}

	ctx.JSON(http.StatusOK, res)
}

type getRouteRequest struct {
	Start string `json:"start"`
	Stop string `json:"stop"`
}

type getRouteResponse struct {
	Path []ResponsePoint `json:"path"`
	Floors []ResponseFloor `json:"floors"`
}

type ResponsePoint struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Floor_areas string `json:"floorArea"`
	Url         string `json:"img"`
	Width       string `json:"width"`
	Height      string `json:"height"`
	Floors 		[]string `json:"floors"`
}

type ResponseFloor struct {
	Id          int `json:"id"`
	Name        string `json:"name"`
	Image	Image `json"image"`
}

type Image struct {
	Url         string `json:"url"`
	Width       int `json:"width"`
	Height      int `json:"height"`
}

package database

type Edge struct {
	Id     int
	Point1 string
	Point2 string
}

type Point struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Floor_areas string `json:"floorArea"`
	Url         string `json:"img"`
	Width       string `json:"width"`
	Height      string `json:"height"`
}

type Floor struct {
	Id          int
	Url         string
	Width       int
	Height      int
	Name        string
	Description string
}

type PointFloor struct {
	PointId string
	FloorId int
}

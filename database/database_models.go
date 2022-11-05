package database

type Edge struct {
	Id     int
	Point1 string
	Point2 string
}

type Point struct {
	Id          string
	Name        string
	Description string
	Floor_areas string
}

type Floor struct {
	Id 			int
	Url			string
	Width		int
	Height		int
	Name 		string
	Description string
}

type PointFloor struct {
	PointId	string
	FloorId	int
}
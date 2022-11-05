package database

type Point_floor struct {
	Floor_id string
}

func (db *data) FloorIdSelectByPointId(id string) (*[]Point_floor, error) {
	var floorId []Point_floor

	err := db.Model(&floorId).Where("point_id = ?", id).Select()
	if err != nil {
		return &floorId, err
	}

	return &floorId, nil
}

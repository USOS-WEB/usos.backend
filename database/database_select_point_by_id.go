package database

func (db *data) PointSelectByID(id string) (*Point, error) {
	var point Point

	err := db.Model(&point).Where("id = ?", id).Select()
	if err != nil {
		return &point, err
	}

	return &point, nil
}

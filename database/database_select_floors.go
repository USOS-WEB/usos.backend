package database

func (db *data) FloorSelect() (*[]Floor, error) {
	var floor []Floor

	err := db.Model(&floor).Select()
	if err != nil {
		return &floor, err
	}

	return &floor, nil
}

package areas

import "database/sql"

func FetchAllAreas(db *sql.DB) ([]Area, error) {
	rows, err := db.Query("SELECT id, name, slug FROM Areas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var areas []Area
	for rows.Next() {
		var area Area
		if err := rows.Scan(&area.ID, &area.Name, &area.Slug); err != nil {
			return nil, err
		}
		areas = append(areas, area)
	}
	return areas, nil
}

func FetchAreaByID(db *sql.DB, id int) (Area, error) {
	var area Area
	err := db.QueryRow("SELECT id, name, slug FROM Areas WHERE id = ?", id).Scan(&area.ID, &area.Name, &area.Slug)
	if err != nil {
		return Area{}, err
	}
	return area, nil
}

package developers

import "database/sql"

func FetchAllDevelopers(db *sql.DB) ([]Developer, error) {
	rows, err := db.Query("SELECT id, name, logo_path FROM Developers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var developers []Developer
	for rows.Next() {
		var dev Developer
		if err := rows.Scan(&dev.ID, &dev.Name, &dev.LogoPath); err != nil {
			return nil, err
		}
		developers = append(developers, dev)
	}
	return developers, nil
}

func FetchDeveloperByID(db *sql.DB, id int) (Developer, error) {
	var dev Developer
	err := db.QueryRow("SELECT id, name, logo_path FROM Developers WHERE id = ?", id).Scan(&dev.ID, &dev.Name, &dev.LogoPath)
	if err != nil {
		return Developer{}, err
	}
	return dev, nil
}

package amenities

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Amenity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetAmenities(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name FROM Amenities")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var amenities []Amenity
		for rows.Next() {
			var a Amenity
			if err := rows.Scan(&a.ID, &a.Name); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			amenities = append(amenities, a)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(amenities)
	}
}

package areas

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Area struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func GetAreas(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, slug FROM Areas")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var areas []Area
		for rows.Next() {
			var area Area
			if err := rows.Scan(&area.ID, &area.Name, &area.Slug); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			areas = append(areas, area)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(areas)
	}
}

func GetAreaByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		var area Area
		err := db.QueryRow("SELECT id, name, slug FROM Areas WHERE id = ?", id).Scan(&area.ID, &area.Name, &area.Slug)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Area not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(area)
	}
}

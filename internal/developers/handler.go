package developers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Developer struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LogoPath string `json:"logo_path"`
}

func GetDevelopers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, logo_path FROM Developers")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var developers []Developer
		for rows.Next() {
			var dev Developer
			if err := rows.Scan(&dev.ID, &dev.Name, &dev.LogoPath); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			developers = append(developers, dev)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(developers)
	}
}

func GetDeveloperByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		var dev Developer
		err := db.QueryRow("SELECT id, name, logo_path FROM Developers WHERE id = ?", id).Scan(&dev.ID, &dev.Name, &dev.LogoPath)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Developer not found", http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dev)
	}
}

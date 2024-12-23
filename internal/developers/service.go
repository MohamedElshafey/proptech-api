package developers

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/developers", GetDevelopers(db)).Methods("GET")
	r.HandleFunc("/developers/{id:[0-9]+}", GetDeveloperByID(db)).Methods("GET")
}

package areas

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/areas", GetAreas(db)).Methods("GET")
	r.HandleFunc("/areas/{id:[0-9]+}", GetAreaByID(db)).Methods("GET")
}

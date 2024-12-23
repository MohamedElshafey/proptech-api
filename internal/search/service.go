package search

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterSearchRoutes(r *mux.Router, db *sql.DB) {
	// Search for properties with filters
	r.HandleFunc("/search/properties", SearchPropertiesHandler(db)).Methods("GET")

	// Search for compounds with filters
	r.HandleFunc("/search/compounds", SearchCompoundsHandler(db)).Methods("GET")
}

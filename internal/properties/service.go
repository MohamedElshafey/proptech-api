package properties

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/properties", GetProperties(db)).Methods("GET")
	r.HandleFunc("/properties/{id:[0-9]+}", GetPropertyByID(db)).Methods("GET")
}

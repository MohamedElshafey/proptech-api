package offers

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/offers", GetOffers(db)).Methods("GET")
	r.HandleFunc("/offers/{id:[0-9]+}", GetOfferByID(db)).Methods("GET")
}

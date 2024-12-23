package main

import (
	"log"
	"net/http"

	"github.com/MohamedElshafey/proptech-api/db"
	"github.com/MohamedElshafey/proptech-api/internal/amenities"
	"github.com/MohamedElshafey/proptech-api/internal/areas"
	"github.com/MohamedElshafey/proptech-api/internal/developers"
	"github.com/MohamedElshafey/proptech-api/internal/offers"
	"github.com/MohamedElshafey/proptech-api/internal/properties"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	dbConn := db.InitDB()
	defer dbConn.Close()

	// Set up router
	r := mux.NewRouter()

	// Register routes
	areas.RegisterRoutes(r, dbConn)
	developers.RegisterRoutes(r, dbConn)
	properties.RegisterRoutes(r, dbConn)
	offers.RegisterRoutes(r, dbConn)
	amenities.RegisterRoutes(r, dbConn)

	// Start server
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}

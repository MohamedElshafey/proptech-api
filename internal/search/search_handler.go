package search

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func SearchPropertiesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract filters from query params
		filters := map[string]interface{}{
			"area_id":       r.URL.Query().Get("area_id"),
			"developer_id":  r.URL.Query().Get("developer_id"),
			"compound_id":   r.URL.Query().Get("compound_id"),
			"min_price":     r.URL.Query().Get("min_price"),
			"max_price":     r.URL.Query().Get("max_price"),
			"min_unit_area": r.URL.Query().Get("min_unit_area"),
			"max_unit_area": r.URL.Query().Get("max_unit_area"),
		}

		results, err := FetchProperties(db, filters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	}
}

func SearchCompoundsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract filters from query parameters
		filters := map[string]interface{}{
			"area_id":      parseIntOrNil(r.URL.Query().Get("area_id")),
			"developer_id": parseIntOrNil(r.URL.Query().Get("developer_id")),
			"livable":      parseBoolOrNil(r.URL.Query().Get("livable")),
		}

		// Fetch compounds using the repository function
		results, err := FetchCompounds(db, filters)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return the results as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	}
}

// Helper function to parse integers or return nil if not present
func parseIntOrNil(value string) interface{} {
	if value == "" {
		return nil
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return nil
	}
	return intValue
}

// Helper function to parse booleans or return nil if not present
func parseBoolOrNil(value string) interface{} {
	if value == "" {
		return nil
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return nil
	}
	return boolValue
}

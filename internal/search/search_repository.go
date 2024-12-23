package search

import (
	"database/sql"
	"fmt"
)

type PropertySearchResult struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Slug          string  `json:"slug"`
	MinPrice      float64 `json:"min_price"`
	MaxPrice      float64 `json:"max_price"`
	Currency      string  `json:"currency"`
	MinUnitArea   int     `json:"min_unit_area"`
	MaxUnitArea   int     `json:"max_unit_area"`
	CompoundName  string  `json:"compound_name"`
	AreaName      string  `json:"area_name"`
	DeveloperName string  `json:"developer_name"`
}

type CompoundSearchResult struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	AreaName      string `json:"area_name"`
	DeveloperName string `json:"developer_name"`
	Livable       bool   `json:"livable"`
}

// FetchProperties searches for properties with filters.
func FetchProperties(db *sql.DB, filters map[string]interface{}) ([]PropertySearchResult, error) {
	// Convert empty strings to nil
	for key, value := range filters {
		if value == "" {
			filters[key] = nil
		}
	}

	query := `
	    SELECT
	        p.id, p.name, p.slug, p.min_price, p.max_price, p.currency,
	        p.min_unit_area, p.max_unit_area, c.name AS compound_name,
	        a.name AS area_name, d.name AS developer_name
	    FROM Properties p
	    JOIN Compounds c ON p.compound_id = c.id
	    JOIN Areas a ON p.area_id = a.id
	    JOIN Developers d ON p.developer_id = d.id
	    WHERE
	        (p.area_id = ? OR ? IS NULL)
	        AND (p.developer_id = ? OR ? IS NULL)
	        AND (p.compound_id = ? OR ? IS NULL)
	        AND (p.min_price >= ? OR ? IS NULL)
	        AND (p.max_price <= ? OR ? IS NULL)
	        AND (p.min_unit_area >= ? OR ? IS NULL)
	        AND (p.max_unit_area <= ? OR ? IS NULL);
	`

	// Debugging: Print the query and filters
	fmt.Println("Executing query:", query)
	fmt.Println("With filters:", filters)

	rows, err := db.Query(query,
		filters["area_id"], filters["area_id"],
		filters["developer_id"], filters["developer_id"],
		filters["compound_id"], filters["compound_id"],
		filters["min_price"], filters["min_price"],
		filters["max_price"], filters["max_price"],
		filters["min_unit_area"], filters["min_unit_area"],
		filters["max_unit_area"], filters["max_unit_area"],
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []PropertySearchResult
	for rows.Next() {
		var r PropertySearchResult
		if err := rows.Scan(
			&r.ID, &r.Name, &r.Slug, &r.MinPrice, &r.MaxPrice, &r.Currency,
			&r.MinUnitArea, &r.MaxUnitArea, &r.CompoundName,
			&r.AreaName, &r.DeveloperName,
		); err != nil {
			return nil, err
		}
		results = append(results, r)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// FetchCompounds searches for compounds with filters.
func FetchCompounds(db *sql.DB, filters map[string]interface{}) ([]CompoundSearchResult, error) {
	// Convert empty strings to nil
	for key, value := range filters {
		if value == "" {
			filters[key] = nil
		}
	}

	query := `
        SELECT 
            c.id, c.name, c.slug, a.name AS area_name, 
            d.name AS developer_name, c.livable
        FROM Compounds c
        JOIN Areas a ON c.area_id = a.id
        JOIN Developers d ON c.developer_id = d.id
        WHERE 
            (c.area_id = ? OR ? IS NULL)
            AND (c.developer_id = ? OR ? IS NULL)
            AND (c.livable = ? OR ? IS NULL);
    `

	// Debugging: Print the query and filters
	fmt.Println("Executing query:", query)
	fmt.Println("With filters:", filters)

	rows, err := db.Query(query,
		filters["area_id"], filters["area_id"],
		filters["developer_id"], filters["developer_id"],
		filters["livable"], filters["livable"],
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []CompoundSearchResult
	for rows.Next() {
		var r CompoundSearchResult
		var livable sql.NullBool
		if err := rows.Scan(
			&r.ID, &r.Name, &r.Slug, &r.AreaName,
			&r.DeveloperName, &livable,
		); err != nil {
			return nil, err
		}
		r.Livable = livable.Bool
		results = append(results, r)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

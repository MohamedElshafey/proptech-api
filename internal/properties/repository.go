package properties

import "database/sql"

type PropertyDetails struct {
	ID                  int     `json:"id"`
	Name                string  `json:"name"`
	Slug                string  `json:"slug"`
	CompoundID          int     `json:"compound_id"`
	AreaID              int     `json:"area_id"`
	DeveloperID         int     `json:"developer_id"`
	MinUnitArea         int     `json:"min_unit_area"`
	MaxUnitArea         int     `json:"max_unit_area"`
	MinPrice            float64 `json:"min_price"`
	MaxPrice            float64 `json:"max_price"`
	Currency            string  `json:"currency"`
	MaxInstallmentYears int     `json:"max_installment_years"`
	MinDownPayment      float64 `json:"min_down_payment"`
	FinishingType       string  `json:"finishing_type"`
}

// FetchAllProperties retrieves all properties from the database.
func FetchAllProperties(db *sql.DB) ([]PropertyDetails, error) {
	rows, err := db.Query(`
        SELECT id, name, slug, compound_id, area_id, developer_id, min_unit_area, max_unit_area, 
               min_price, max_price, currency, max_installment_years, min_down_payment, finishing_type
        FROM Properties
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var properties []PropertyDetails
	for rows.Next() {
		var p PropertyDetails
		if err := rows.Scan(
			&p.ID, &p.Name, &p.Slug, &p.CompoundID, &p.AreaID, &p.DeveloperID,
			&p.MinUnitArea, &p.MaxUnitArea, &p.MinPrice, &p.MaxPrice, &p.Currency,
			&p.MaxInstallmentYears, &p.MinDownPayment, &p.FinishingType,
		); err != nil {
			return nil, err
		}
		properties = append(properties, p)
	}
	return properties, nil
}

// FetchPropertyByID retrieves a single property by its ID.
func FetchPropertyByID(db *sql.DB, id int) (PropertyDetails, error) {
	var p PropertyDetails
	err := db.QueryRow(`
        SELECT id, name, slug, compound_id, area_id, developer_id, min_unit_area, max_unit_area, 
               min_price, max_price, currency, max_installment_years, min_down_payment, finishing_type
        FROM Properties
        WHERE id = ?
    `, id).Scan(
		&p.ID, &p.Name, &p.Slug, &p.CompoundID, &p.AreaID, &p.DeveloperID,
		&p.MinUnitArea, &p.MaxUnitArea, &p.MinPrice, &p.MaxPrice, &p.Currency,
		&p.MaxInstallmentYears, &p.MinDownPayment, &p.FinishingType,
	)
	if err != nil {
		return PropertyDetails{}, err
	}
	return p, nil
}

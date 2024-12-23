package offers

import "database/sql"

type Offer struct {
	ID                       int      `json:"id"`
	CompoundID               int      `json:"compound_id"`
	Description              string   `json:"description"`
	Discount                 *float64 `json:"discount,omitempty"`
	MaintenancePercent       *float64 `json:"maintenance_percent,omitempty"`
	DeliveryPaymentPercent   *float64 `json:"delivery_payment_percent,omitempty"`
	EqualInstallmentsPercent *float64 `json:"equal_installments_percent,omitempty"`
	DownPaymentPercent       *float64 `json:"down_payment_percent,omitempty"`
	Years                    int      `json:"years"`
}

// FetchAllOffers retrieves all offers from the database.
func FetchAllOffers(db *sql.DB) ([]Offer, error) {
	rows, err := db.Query(`
        SELECT id, compound_id, description, discount, maintenance_percent, delivery_payment_percent, 
               equal_installments_percent, down_payment_percent, years
        FROM Offers
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offers []Offer
	for rows.Next() {
		var o Offer
		if err := rows.Scan(
			&o.ID, &o.CompoundID, &o.Description, &o.Discount, &o.MaintenancePercent,
			&o.DeliveryPaymentPercent, &o.EqualInstallmentsPercent, &o.DownPaymentPercent, &o.Years,
		); err != nil {
			return nil, err
		}
		offers = append(offers, o)
	}
	return offers, nil
}

// FetchOfferByID retrieves a single offer by its ID.
func FetchOfferByID(db *sql.DB, id int) (Offer, error) {
	var o Offer
	err := db.QueryRow(`
        SELECT id, compound_id, description, discount, maintenance_percent, delivery_payment_percent, 
               equal_installments_percent, down_payment_percent, years
        FROM Offers
        WHERE id = ?
    `, id).Scan(
		&o.ID, &o.CompoundID, &o.Description, &o.Discount, &o.MaintenancePercent,
		&o.DeliveryPaymentPercent, &o.EqualInstallmentsPercent, &o.DownPaymentPercent, &o.Years,
	)
	if err != nil {
		return Offer{}, err
	}
	return o, nil
}

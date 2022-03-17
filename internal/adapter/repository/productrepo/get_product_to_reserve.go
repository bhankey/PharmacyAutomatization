package productrepo

import (
	"context"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

// TODO refactor. Don't like way of working with products.
func (r *Repository) GetProductToReserve(
	ctx context.Context,
	productName string,
	position string,
) (entities.ProductItem, error) {
	errBase := fmt.Sprintf("productrepo.GetProductToReserve(%s, %s)", productName, position)

	const query = `
		SELECT id,
		       product_id,
		       receipt_id,
		       pharmacy_id,
		       position,
		       manufactured_time,
		       reservation,
		       is_sold, 
		       is_expired,
		       priority
		FROM product_item INNER JOIN product ON product_item.product_id = product.id
		WHERE product.name = $1 	
		  AND product_item.position = $2
		  AND product_item.is_expired = false
		  AND product_item.is_sold = false
		  AND product_item.reservation ISNULL
		  AND product_item.manufactured_time + product.expiration_date >= NOW()
		ORDER BY product_item.priority
		LIMIT 1
`

	row := struct {
		ID               int    `db:"id"`
		ProductID        int    `db:"product_id"`
		PharmacyID       int    `db:"pharmacy_id"`
		ReceiptID        int    `db:"receipt_id"`
		Position         string `db:"position"`
		ManufacturedTime string `db:"manufactured_time"`
		ReservationUUID  string `db:"reservation_uuid"`
		IsSold           bool   `db:"is_sold"`
		IsExpired        bool   `db:"is_expired"`
		Priority         int    `db:"priority"`
	}{}

	if err := r.master.GetContext(ctx, &row, query, productName, position); err != nil {
		return entities.ProductItem{}, fmt.Errorf("%s: QueryError: %w", errBase, err)
	}

	return entities.ProductItem{
		ID:               row.ID,
		ProductID:        row.ProductID,
		PharmacyID:       row.PharmacyID,
		ReceiptID:        row.ReceiptID,
		Position:         row.Position,
		ManufacturedTime: row.ManufacturedTime,
		ReservationUUID:  row.ReservationUUID,
		IsSold:           row.IsSold,
		IsExpired:        row.IsExpired,
		Priority:         row.Priority,
	}, nil
}

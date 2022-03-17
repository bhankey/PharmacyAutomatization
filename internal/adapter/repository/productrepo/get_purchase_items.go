package productrepo

import (
	"context"
	"fmt"
	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (r *Repository) GetPurchaseProducts(ctx context.Context, pharmacyID int, purchaseUUID string) ([]entities.PurchaseProductItem, error) {
	errBase := fmt.Sprintf("pharmacyrepo.GetAvailablePharmacyProducts(%d)", pharmacyID)

	const query = `
		SELECT p.name, p.price, COUNT(p_item.id) as count
		FROM product_item p_item
		INNER JOIN product p on p.id = p_item.product_id
		WHERE pharmacy_id = $1
		  AND p_item.reservation = $2
		GROUP BY p.id
		ORDER BY p_item.priority
`

	type row struct {
		Name  string `db:"name"`
		Price int    `db:"price"`
		Count int    `db:"count"`
	}

	rows := make([]row, 0)
	if err := r.slave.SelectContext(ctx, &rows, query, pharmacyID, purchaseUUID); err != nil {
		return nil, fmt.Errorf("%s: QueryError: %w", errBase, err)
	}

	result := make([]entities.PurchaseProductItem, 0, len(rows))
	for _, row := range rows {
		result = append(result, entities.PurchaseProductItem{
			Name:  row.Name,
			Price: row.Price,
			Count: row.Count,
		})
	}

	return result, nil
}

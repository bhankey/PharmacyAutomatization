package receiptrepo

import (
	"context"
	"fmt"
)

func (r *Repository) CreateReceipt(ctx context.Context, userID, pharmacyID, sum, discount int, purchaseUUID string) (int, error) {
	errBase := fmt.Sprintf("receiptrepo.CreateReceipt(%d, %d, %d, %d)", userID, pharmacyID, sum, discount)

	const query string = `
		INSERT INTO receipt(user_id, pharmacy_id, sum, discount, purchase_uuid)
					VALUES ($1, $2, $3, $4, $5)
`

	res, err := r.master.ExecContext(
		ctx,
		query,
		userID,
		pharmacyID,
		sum,
		discount,
		purchaseUUID,
	)
	if err != nil {
		return 0, fmt.Errorf("%s: QueryError: %w", errBase, err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: LastInsertID.Error: %w", errBase, err)
	}

	return int(ID), nil
}

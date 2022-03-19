package productrepo

import (
	"context"
	"fmt"
)

func (r *Repository) SetProductItemsSold(ctx context.Context, receiptID int, purchaseUUID string) error {
	errBase := fmt.Sprintf("productrepo.SetProductSold(%s)", purchaseUUID)

	const query = `
		UPDATE product_item 
		SET is_sold = true, receipt_id = $1
		WHERE reservation = $2
`
	if _, err := r.master.ExecContext(
		ctx,
		query,
		receiptID,
		purchaseUUID,
	); err != nil {
		return fmt.Errorf("%s: QueryError: %w", errBase, err)
	}

	return nil
}

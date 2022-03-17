package productrepo

import (
	"context"
	"fmt"
)

func (r *Repository) Reserve(ctx context.Context, productID int, purchaseUUID string) error {
	errBase := fmt.Sprintf("productrepo.Reserve(%d, %s)", productID, purchaseUUID)

	const query = `
		UPDATE product_item 
		SET reservation = $1
		WHERE id = $2
		LIMIT 1
`
	if _, err := r.master.ExecContext(
		ctx,
		query,
		purchaseUUID,
		productID,
	); err != nil {
		return fmt.Errorf("%s: failed to update password: %w", errBase, err)
	}

	return nil
}

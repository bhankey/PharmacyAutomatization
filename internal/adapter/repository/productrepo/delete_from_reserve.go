package productrepo

import (
	"context"
	"fmt"
)

func (r *Repository) DeleteFromReservation(ctx context.Context, productID int) error {
	errBase := fmt.Sprintf("productrepo.Reserve(%d, %s)", productID)

	const query = `
		UPDATE product_item 
		SET reservation = NULL
		WHERE id = $1
`
	if _, err := r.master.ExecContext(
		ctx,
		query,
		productID,
	); err != nil {
		return fmt.Errorf("%s: failed to update password: %w", errBase, err)
	}

	return nil
}

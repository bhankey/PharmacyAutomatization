package tokenrepo

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func (r *TokenRepo) DeactivateTokenByIDs(ctx context.Context, tokenIDs []int) error {
	const query = `
		UPDATE refresh_tokens 
		SET is_available = false
		WHERE id IN (?)
`

	resultingQuery, params, err := sqlx.In(query, tokenIDs)
	if err != nil {
		return fmt.Errorf("failed to prepare query with sqlx.In error: %w", err)
	}

	resultingQuery = r.master.Rebind(resultingQuery)
	if _, err := r.master.ExecContext(ctx, resultingQuery, params...); err != nil {
		return fmt.Errorf("failed to update refresh tokens error: %w", err)
	}

	return nil
}

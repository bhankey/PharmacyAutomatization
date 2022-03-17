package userrepo

import (
	"context"
	"fmt"
)

func (r *Repository) UpdatePassword(ctx context.Context, email string, newPasswordHash string) error {
	errBase := fmt.Sprintf("userrepo.UpdatePassword(%s, %s)", email, newPasswordHash)

	const query = `
		UPDATE users 
		SET password_hash = $1
		WHERE email = $2
`
	if _, err := r.master.ExecContext(
		ctx,
		query,
		newPasswordHash,
		email,
	); err != nil {
		return fmt.Errorf("%s: failed to update password: %w", errBase, err)
	}

	return nil
}

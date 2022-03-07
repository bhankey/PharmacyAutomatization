package userrepo

import (
	"context"
)

func (r *UserRepo) UpdatePassword(ctx context.Context, email string, newPasswordHash string) error {
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
		return err
	}

	return nil
}

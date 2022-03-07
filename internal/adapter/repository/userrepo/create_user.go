package userrepo

import (
	"context"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (r *UserRepo) CreateUser(ctx context.Context, user entities.User) error {
	const query = `
		INSERT INTO users(name, surname, email, password_hash, default_pharmacy_id)
							VALUE ($1, $2, $3, $4, $5)
`

	if _, err := r.master.ExecContext(
		ctx,
		query,
		user.Name,
		user.Surname,
		user.PasswordHash,
		user.DefaultPharmacyID,
	); err != nil {
		return fmt.Errorf("failed to insert refresh token: %w", err)
	}

	return nil
}

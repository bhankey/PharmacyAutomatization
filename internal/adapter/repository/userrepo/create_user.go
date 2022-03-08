package userrepo

import (
	"context"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (r *UserRepo) CreateUser(ctx context.Context, user entities.User) error {
	errBase := fmt.Sprintf("userrepo.CreateUser(%v)", user)

	const query = `
		INSERT INTO users(name, surname, email, password_hash)
							VALUES ($1, $2, $3, $4)
`

	if _, err := r.master.ExecContext(
		ctx,
		query,
		user.Name,
		user.Surname,
		user.Email,
		user.PasswordHash,
	); err != nil {
		return fmt.Errorf("%s: failed to create user: %w", errBase, err)
	}

	return nil
}

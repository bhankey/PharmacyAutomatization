package userrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/bhankey/BD_lab/backend/internal/entities"
	"github.com/bhankey/BD_lab/backend/internal/repository"
)

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	const query = `
		SELECT id, name, surname, email, password_hash, default_pharmacy_id
		FROM users
		WHERE email = $1
		LIMIT 1
`

	row := struct {
		ID                int           `db:"id"`
		Name              string        `db:"name"`
		Email             string        `db:"email"`
		PasswordHash      string        `db:"password_hash"`
		DefaultPharmacyID sql.NullInt64 `db:"default_pharmacy_id"`
	}{}

	var password string
	if err := r.slave.GetContext(ctx, &password, query, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, repository.ErrNoEntity
		}

		return entities.User{}, fmt.Errorf("failed to get user password_hash error: %w", err)
	}

	return entities.User{
		ID:                row.ID,
		Name:              row.Name,
		Email:             row.Email,
		PasswordHash:      row.PasswordHash,
		DefaultPharmacyID: int(row.DefaultPharmacyID.Int64),
	}, nil
}

// nolint: dupl, nolintlint
package userrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/apperror"
	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	errBase := fmt.Sprintf("userrepo.GetUserByEmail(%s)", email)

	const query = `
		SELECT id, name, surname, email, password_hash, role, default_pharmacy_id
		FROM users
		WHERE email = $1
		LIMIT 1
`

	row := struct {
		ID                int           `db:"id"`
		Name              string        `db:"name"`
		Surname           string        `json:"surname"`
		Email             string        `db:"email"`
		Role              string        `db:"role"`
		PasswordHash      string        `db:"password_hash"`
		DefaultPharmacyID sql.NullInt64 `db:"default_pharmacy_id"`
	}{}

	if err := r.slave.GetContext(ctx, &row, query, email); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, apperror.ErrNoEntity
		}

		return entities.User{}, fmt.Errorf("%s: failed to get user by email error: %w", errBase, err)
	}

	return entities.User{
		ID:                row.ID,
		Name:              row.Name,
		Surname:           row.Surname,
		Email:             row.Email,
		PasswordHash:      row.PasswordHash,
		Role:              entities.Role(row.Role),
		DefaultPharmacyID: int(row.DefaultPharmacyID.Int64),
	}, nil
}

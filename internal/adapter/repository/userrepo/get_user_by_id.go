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

func (r *Repository) GetUserByID(ctx context.Context, id int) (entities.User, error) {
	errBase := fmt.Sprintf("userrepo.GetUserByID(%d)", id)

	const query = `
		SELECT id, name, surname, email, password_hash, role, default_pharmacy_id
		FROM users
		WHERE id = $1
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

	if err := r.slave.GetContext(ctx, &row, query, id); err != nil {
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
		DefaultPharmacyID: int(row.DefaultPharmacyID.Int64),
	}, nil
}

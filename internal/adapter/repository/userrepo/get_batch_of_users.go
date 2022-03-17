package userrepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (r *Repository) GetBatchOfUsers(ctx context.Context, lastClientID int, limit int) ([]entities.User, error) {
	errBase := fmt.Sprintf("userrepo.GetBatchOfUsers(%d, %d)", lastClientID, limit)

	const query = `
		SELECT id, name, surname, email, role, password_hash, default_pharmacy_id
		FROM users
		WHERE id > $1
		LIMIT $2
`

	type row struct {
		ID                int           `db:"id"`
		Name              string        `db:"name"`
		Surname           string        `json:"surname"`
		Email             string        `db:"email"`
		PasswordHash      string        `db:"password_hash"`
		Role              string        `db:"role"`
		DefaultPharmacyID sql.NullInt64 `db:"default_pharmacy_id"`
	}

	rows := make([]row, 0)
	if err := r.slave.SelectContext(ctx, &rows, query, lastClientID, limit); err != nil {
		return nil, fmt.Errorf("%s: failed to get user by email error: %w", errBase, err)
	}

	result := make([]entities.User, 0, len(rows))

	for _, row := range rows {
		result = append(result, entities.User{
			ID:                row.ID,
			Name:              row.Name,
			Surname:           row.Surname,
			Email:             row.Email,
			Password:          row.PasswordHash,
			PasswordHash:      row.PasswordHash,
			Role:              entities.Role(row.Role),
			DefaultPharmacyID: int(row.DefaultPharmacyID.Int64),
		})
	}

	return result, nil
}

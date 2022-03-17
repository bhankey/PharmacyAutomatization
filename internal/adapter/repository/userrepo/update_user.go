package userrepo

import (
	"context"
	"fmt"
	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (r *Repository) UpdateUser(ctx context.Context, user entities.User) error {
	errBase := fmt.Sprintf("userrepo.UpdateUser(%v)", user)

	const query = `
		UPDATE users SET
			name = COALESCE($1, name),
			surname = COALESCE($2, surname),
			email = COALESCE($3, email),
		    role = COALESCE($4, role),
		    default_pharmacy_id = COALESCE($5, default_pharmacy_id)
		WHERE id = $6
`

	row := struct {
		Name              *string `db:"name"`
		Surname           *string `json:"surname"`
		Email             *string `db:"email"`
		Role              *string `db:"role"`
		DefaultPharmacyID *int    `db:"default_pharmacy_id"`
	}{}

	if user.Name != "" {
		row.Name = &user.Name
	}
	if user.Surname != "" {
		row.Surname = &user.Surname
	}
	if user.Email != "" {
		row.Email = &user.Email
	}
	if user.Role != "" {
		row.Role = (*string)(&user.Role)
	}
	if user.DefaultPharmacyID != 0 {
		row.DefaultPharmacyID = &user.DefaultPharmacyID
	}

	if _, err := r.master.ExecContext(
		ctx,
		query,
		row.Name, row.Surname, row.Email, row.Role, row.DefaultPharmacyID, user.ID,
	); err != nil {
		return fmt.Errorf("%s: failed to update user: %w", errBase, err)
	}

	return nil
}

package addressrepo

import (
	"context"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (r *Repository) CreateAddress(ctx context.Context, address entities.Address) (int, error) {
	errBase := fmt.Sprintf("addresses.CreateAddress(%v)", address)

	const query string = `
		INSERT INTO addresses(city, street, house)
					VALUES ($1, $2, $3)
`

	res, err := r.master.ExecContext(
		ctx,
		query,
		address.City,
		address.Street,
		address.House,
	)
	if err != nil {
		return 0, fmt.Errorf("%s: QueryError: %w", errBase, err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: LastInsertID.Error: %w", errBase, err)
	}

	return int(ID), nil
}

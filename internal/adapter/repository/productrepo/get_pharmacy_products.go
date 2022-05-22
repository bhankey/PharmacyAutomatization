package productrepo

import (
	"context"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (r *Repository) GetAvailablePharmacyProducts(
	ctx context.Context,
	pharmacyID int,
) ([]entities.PharmacyProductItem, error) {
	errBase := fmt.Sprintf("pharmacyrepo.GetAvailablePharmacyProducts(%d)", pharmacyID)

	const query = `
		SELECT p.name,
		       p.price,
		       p.instruction_url,
		       p.img_url,
		       p.comment,
		       p.recipe_only,
		       p_item.position,
		       COUNT(p_item.id) as count
		FROM product_item p_item
		INNER JOIN product p on p.id = p_item.product_id
		WHERE pharmacy_id = $1
		  AND p_item.is_expired = false
		  AND p_item.is_sold = false
		  AND p_item.reservation ISNULL
		  AND p_item.manufactured_time + p.expiration_date >= NOW()
		GROUP BY p.id, p_item.position, p_item.priority
		HAVING count(p_item.id) > 0
		ORDER BY p_item.priority
`

	type row struct {
		Name           string `db:"name"`
		Price          int    `db:"price"`
		InstructionURL string `db:"instruction_url"`
		ImgURL         string `db:"img_url"`
		Comment        string `db:"comment"`
		RecipeOnly     bool   `db:"recipe_only"`
		Position       string `db:"position"`
		Count          int    `db:"count"`
	}

	rows := make([]row, 0)
	if err := r.slave.SelectContext(ctx, &rows, query, pharmacyID); err != nil {
		return nil, fmt.Errorf("%s: QueryError: %w", errBase, err)
	}

	result := make([]entities.PharmacyProductItem, 0, len(rows))
	for _, row := range rows {
		result = append(result, entities.PharmacyProductItem{
			Name:           row.Name,
			Price:          row.Price,
			InstructionURL: row.InstructionURL,
			ImgURL:         row.ImgURL,
			Comment:        row.Comment,
			RecipeOnly:     row.RecipeOnly,
			Position:       row.Position,
			Count:          row.Count,
		})
	}

	return result, nil
}

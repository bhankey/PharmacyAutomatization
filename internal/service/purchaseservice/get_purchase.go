package purchaseservice

import (
	"context"
	"fmt"
	"github.com/bhankey/pharmacy-automatization/internal/entities"
	"math"
)

func (s *Service) GetPurchase(ctx context.Context, pharmacyID int, purchaseUUID string, isSocialCard bool) (entities.Purchase, error) {
	errBase := fmt.Sprintf("purchaseservice.GetPurchase(%d, %s)", pharmacyID, purchaseUUID)

	products, err := s.productRepo.GetPurchaseProducts(ctx, pharmacyID, purchaseUUID)
	if err != nil {
		return entities.Purchase{}, fmt.Errorf("%s: failed to get purchase: %w", errBase, err)
	}

	sum := 0
	for _, product := range products {
		sum += product.Price
	}

	if isSocialCard {
		sum = int(math.Floor(float64(sum) * 0.95))
	}

	return entities.Purchase{
		Price:    sum,
		Products: products,
	}, nil
}

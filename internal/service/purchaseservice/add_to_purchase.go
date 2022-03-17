package purchaseservice

import (
	"context"
	"fmt"
)

func (s *Service) AddToPurchase(ctx context.Context, productName string, position string, purchaseUUID string) error {
	errBase := fmt.Sprintf("purchaseservice.AddToPurchase(%s, %s, %s)", productName, position, purchaseUUID)

	product, err := s.productRepo.GetProductToReserve(ctx, productName, position)
	if err != nil {
		return fmt.Errorf("%s: Failed to get product item: %w", errBase, err)
	}

	if err := s.productRepo.Reserve(ctx, product.ID, purchaseUUID); err != nil {
		return fmt.Errorf("%s: failed to reserve product: %w", errBase, err)
	}

	return nil
}

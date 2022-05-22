package purchaseservice

import (
	"context"
	"fmt"
)

func (s *Service) DeleteFromPurchase(ctx context.Context, productName string, position string, purchaseUUID string) error {
	errBase := fmt.Sprintf("purchaseservice.AddToPurchase(%s, %s, %s)", productName, position, purchaseUUID)

	product, err := s.productRepo.GetReservedProduct(ctx, productName, position, purchaseUUID)
	if err != nil {
		return fmt.Errorf("%s: Failed to get product item: %w", errBase, err)
	}

	if err := s.productRepo.DeleteFromReservation(ctx, product.ID); err != nil {
		return fmt.Errorf("%s: failed to reserve product: %w", errBase, err)
	}

	return nil
}

package purchaseservice

import (
	"context"
	"fmt"
	"math"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (s *Service) ConfirmPurchase(
	ctx context.Context,
	userID, pharmacyID int,
	purchaseUUID string,
	isSocialCardUsed bool,
) error {
	errBase := fmt.Sprintf("purchaseservice.ConfirmPurchase(%s, %v)", purchaseUUID, isSocialCardUsed)

	products, err := s.productRepo.GetPurchaseProducts(ctx, pharmacyID, purchaseUUID)
	if err != nil {
		return fmt.Errorf("%s: failed to get purchase: %w", errBase, err)
	}

	sum := 0
	for _, product := range products {
		sum += product.Price
	}

	discount := 0
	if isSocialCardUsed {
		discount = int(math.Floor(float64(sum) * (1 - entities.SocialCardDiscountMultiplier)))
	}

	receiptID, err := s.receiptRepo.CreateReceipt(ctx, userID, pharmacyID, sum, discount, purchaseUUID)
	if err != nil {
		return fmt.Errorf("%s: failed to create receipt: %w", errBase, err)
	}

	// TODO handle this errors by cron
	if err := s.productRepo.SetProductItemsSold(ctx, receiptID, purchaseUUID); err != nil {
		return fmt.Errorf("%s: failed to set product as sold: %w", errBase, err)
	}

	return nil
}

package pharmacyservice

import (
	"context"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

type Service struct {
	pharmacyRepo PharmacyRepo
}

type PharmacyRepo interface {
	CreatePharmacy(ctx context.Context, pharmacy entities.Pharmacy) error
	GetBatchOfPharmacies(ctx context.Context, lastPharmacyID int, limit int) ([]entities.Pharmacy, error)
	GetAvailablePharmacyProducts(ctx context.Context, pharmacyID int) ([]entities.PharmacyProductItem, error)
}

func NewPharmacyService(
	pharmacyRepo PharmacyRepo,
) *Service {
	return &Service{
		pharmacyRepo: pharmacyRepo,
	}
}

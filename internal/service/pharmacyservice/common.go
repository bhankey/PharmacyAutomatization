package pharmacyservice

import (
	"context"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (s *Service) CreatePharmacy(ctx context.Context, pharmacy entities.Pharmacy) error {
	var err error

	pharmacy.Address.ID, err = s.addressRepo.CreateAddress(ctx, pharmacy.Address)
	if err != nil {
		return err
	}

	return s.pharmacyRepo.CreatePharmacy(ctx, pharmacy)
}

func (s *Service) GetBatchOfPharmacies(ctx context.Context, lastPharmacyID int, limit int) ([]entities.Pharmacy, error) {
	return s.pharmacyRepo.GetBatchOfPharmacies(ctx, lastPharmacyID, limit)
}

func (s *Service) GetPharmacyProducts(ctx context.Context, pharmacyID int) ([]entities.PharmacyProductItem, error) {
	return s.pharmacyRepo.GetAvailablePharmacyProducts(ctx, pharmacyID)
}

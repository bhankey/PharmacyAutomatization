package purchaseservice

import (
	"context"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

type Service struct {
	productRepo ProductRepo
	receiptRepo ReceiptRepo
}

type ReceiptRepo interface {
	CreateReceipt(ctx context.Context, userID, pharmacyID, sum, discount int) (int, error)
}

type ProductRepo interface {
	GetProductToReserve(ctx context.Context, productName string, position string) (entities.ProductItem, error)
	Reserve(ctx context.Context, productID int, purchaseUUID string) error
	GetPurchaseProducts(ctx context.Context, pharmacyID int, purchaseUUID string) ([]entities.PurchaseProductItem, error)
	SetProductItemsSold(ctx context.Context, receiptID int, purchaseUUID string) error
}

func NewPurchaseService(
	productRepo ProductRepo,
	receiptRepo ReceiptRepo,
) *Service {
	return &Service{
		productRepo: productRepo,
		receiptRepo: receiptRepo,
	}
}

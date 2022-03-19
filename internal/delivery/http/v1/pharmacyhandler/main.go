package pharmacyhandler

import (
	"context"

	deliveryhttp "github.com/bhankey/pharmacy-automatization/internal/delivery/http"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/middleware"
	"github.com/bhankey/pharmacy-automatization/internal/entities"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Router chi.Router

	pharmacySrv PharmacySrv

	*deliveryhttp.BaseHandler
}

type PharmacySrv interface {
	CreatePharmacy(ctx context.Context, pharmacy entities.Pharmacy) error
	GetBatchOfPharmacies(ctx context.Context, lastPharmacyID int, limit int) ([]entities.Pharmacy, error)
	GetPharmacyProducts(ctx context.Context, pharmacyID int) ([]entities.PharmacyProductItem, error)
}

func NewPharmacyHandler(
	baseHandler *deliveryhttp.BaseHandler,
	pharmacySrv PharmacySrv,
	authMiddleware *middleware.AuthMiddleware,
) *Handler {
	router := chi.NewRouter()

	handler := &Handler{
		Router:      router,
		pharmacySrv: pharmacySrv,
		BaseHandler: baseHandler,
	}

	handler.initRoutes(router, authMiddleware)

	return handler
}

func (h *Handler) initRoutes(router chi.Router, authMiddleware *middleware.AuthMiddleware) {
	router.Use(authMiddleware.Middleware)
	router.Post("/create", h.create)
	router.Get("/all", h.all)
	router.Get("/products", h.products)
}

package purchasehandler

import (
	"context"

	deliveryhttp "github.com/bhankey/pharmacy-automatization/internal/delivery/http"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/middleware"
	"github.com/bhankey/pharmacy-automatization/internal/entities"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Router chi.Router

	purchaseSrv PurchaseSrv

	*deliveryhttp.BaseHandler
}

type PurchaseSrv interface {
	AddToPurchase(ctx context.Context, productName string, position string, purchaseUUID string) error
	ConfirmPurchase(ctx context.Context, userID, pharmacyID int, purchaseUUID string, isSocialCardUsed bool) error
	GetPurchase(ctx context.Context, pharmacyID int, purchaseUUID string, isSocialCard bool) (entities.Purchase, error)
}

func NewPurchaseHandler(
	baseHandler *deliveryhttp.BaseHandler,
	purchaseSrv PurchaseSrv,
	authMiddleware *middleware.AuthMiddleware,
) *Handler {
	router := chi.NewRouter()

	handler := &Handler{
		Router:      router,
		purchaseSrv: purchaseSrv,
		BaseHandler: baseHandler,
	}

	handler.initRoutes(router, authMiddleware)

	return handler
}

func (h *Handler) initRoutes(router chi.Router, authMiddleware *middleware.AuthMiddleware) {
	router.Use(authMiddleware.Middleware)

	router.Post("/add", h.add)
	router.Post("/confirm", h.confirm)
	router.Get("/show", h.show)
}

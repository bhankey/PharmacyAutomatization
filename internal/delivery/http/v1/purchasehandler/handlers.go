package purchasehandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bhankey/pharmacy-automatization/internal/apperror"
	deliveryhttp "github.com/bhankey/pharmacy-automatization/internal/delivery/http"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/v1/models"
	"github.com/bhankey/pharmacy-automatization/internal/entities"
	"github.com/go-openapi/strfmt"
)

func (h *Handler) add(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer func() { _ = r.Body.Close() }()
	var req models.PurchaseAddRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)
	if err != nil {
		h.WriteErrorResponse(ctx, w, apperror.NewClientError(apperror.WrongRequest, err))

		return
	}

	if err := req.Validate(strfmt.NewFormats()); err != nil {
		h.WriteErrorResponse(ctx, w, apperror.NewClientError(apperror.WrongRequest, err))

		return
	}

	if err := h.purchaseSrv.AddToPurchase(ctx, req.ProductName, req.Position, req.PurchaseUUID); err != nil {
		h.WriteErrorResponse(ctx, w, err)

		return
	}

	deliveryhttp.WriteResponse(w, models.BaseResponse{
		Error:   "",
		Success: true,
	})
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer func() { _ = r.Body.Close() }()
	var req models.PurchaseAddRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)
	if err != nil {
		h.WriteErrorResponse(ctx, w, apperror.NewClientError(apperror.WrongRequest, err))

		return
	}

	if err := req.Validate(strfmt.NewFormats()); err != nil {
		h.WriteErrorResponse(ctx, w, apperror.NewClientError(apperror.WrongRequest, err))

		return
	}

	if err := h.purchaseSrv.DeleteFromPurchase(ctx, req.ProductName, req.Position, req.PurchaseUUID); err != nil {
		h.WriteErrorResponse(ctx, w, err)

		return
	}

	deliveryhttp.WriteResponse(w, models.BaseResponse{
		Error:   "",
		Success: true,
	})
}

func (h *Handler) show(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	query := r.URL.Query()

	purchaseUUID := query.Get("purchase_uuid")
	if purchaseUUID == "" {
		h.WriteErrorResponse(
			ctx,
			w,
			apperror.NewClientError(
				apperror.WrongRequest, fmt.Errorf("wrong requset")), // nolint: goerr113
		)

		return
	}

	isSocialCard := query.Get("is_social_card")
	pharmacyID, _ := ctx.Value(entities.PharmacyID).(int)

	purchase, err := h.purchaseSrv.GetPurchase(ctx, pharmacyID, purchaseUUID, isSocialCard == "true")
	if err != nil {
		h.WriteErrorResponse(ctx, w, err)

		return
	}

	items := make([]*models.PurchaseShowResponseItemsItems0, 0, len(purchase.Products))

	for _, product := range purchase.Products {
		items = append(items, &models.PurchaseShowResponseItemsItems0{
			Count: int64(product.Count),
			Name:  product.Name,
			Price: int64(product.Price),
		})
	}
	deliveryhttp.WriteResponse(w, models.PurchaseShowResponse{
		Items: items,
		Price: int64(purchase.Price),
	})
}

func (h *Handler) confirm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer func() { _ = r.Body.Close() }()
	var req models.PurchaseConfirmRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)
	if err != nil {
		h.WriteErrorResponse(ctx, w, apperror.NewClientError(apperror.WrongRequest, err))

		return
	}

	if err := req.Validate(strfmt.NewFormats()); err != nil {
		h.WriteErrorResponse(ctx, w, apperror.NewClientError(apperror.WrongRequest, err))

		return
	}

	pharmacyID, _ := ctx.Value(entities.PharmacyID).(int)
	userID, _ := ctx.Value(entities.UserID).(int)

	if pharmacyID <= 0 {
		h.WriteErrorResponse(
			ctx,
			w,
			apperror.NewClientError(
				apperror.WrongRequest,
				fmt.Errorf("failed to get user pharmacy or id")), // nolint: goerr113
		)
	}

	err = h.purchaseSrv.ConfirmPurchase(ctx, userID, pharmacyID, string(req.PurchaseUUID), req.IsSocialCard)
	if err != nil {
		h.WriteErrorResponse(ctx, w, err)

		return
	}

	deliveryhttp.WriteResponse(w, models.BaseResponse{
		Error:   "",
		Success: true,
	})
}

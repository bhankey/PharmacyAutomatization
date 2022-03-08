package userhandler

import (
	"encoding/json"
	"net/http"

	"github.com/bhankey/pharmacy-automatization/internal/apperror"
	deliveryhttp "github.com/bhankey/pharmacy-automatization/internal/delivery/http"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/v1/models"
	"github.com/bhankey/pharmacy-automatization/internal/entities"
	"github.com/go-openapi/strfmt"
)

func (h *UserHandler) register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer func() { _ = r.Body.Close() }()
	var req models.RegisterRequest

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

	user := entities.User{
		ID:       0,
		Name:     req.Name,
		Surname:  req.Surname,
		Email:    req.Email.String(),
		Password: *req.Password,
	}
	if err := h.userSrv.Registry(ctx, user); err != nil {
		h.WriteErrorResponse(ctx, w, err)

		return
	}

	deliveryhttp.WriteResponse(w, models.BaseResponse{
		Error:   "",
		Success: true,
	})
}

func (h *UserHandler) requestToChangePassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer func() { _ = r.Body.Close() }()
	var req models.RequestPasswordChangeRequest

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

	if err := h.userSrv.RequestToResetPassword(ctx, req.Email.String()); err != nil {
		h.WriteErrorResponse(ctx, w, err)

		return
	}

	deliveryhttp.WriteResponse(w, models.BaseResponse{
		Error:   "",
		Success: true,
	})
}

func (h *UserHandler) changePassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer func() { _ = r.Body.Close() }()
	var req models.PasswordChangeRequest

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

	if err := h.userSrv.ResetPassword(ctx, req.Email.String(), *req.Code, *req.NewPassword); err != nil {
		h.WriteErrorResponse(ctx, w, err)

		return
	}

	deliveryhttp.WriteResponse(w, models.BaseResponse{
		Error:   "",
		Success: true,
	})
}

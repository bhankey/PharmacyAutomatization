package userhandler

import (
	"encoding/json"
	"net/http"

	deliveryhttp "github.com/bhankey/pharmacy-automatization/internal/delivery/http"
	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/v1/models"
	"github.com/go-openapi/strfmt"
)

func (h *UserHandler) login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer func() { _ = r.Body.Close() }()
	var req models.UserLoginRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)
	if err != nil {
		h.WriteErrorResponse(ctx, w, err, true)

		return
	}

	if err := req.Validate(strfmt.NewFormats()); err != nil {
		h.WriteErrorResponse(ctx, w, err, true)

		return
	}

	token, err := h.userSrv.Login(ctx, req.Email, req.Password)
	if err != nil {
		h.WriteErrorResponse(ctx, w, err, false)

		return
	}

	deliveryhttp.WriteResponse(w, models.UserLoginResponse{
		Token: token,
	})
}

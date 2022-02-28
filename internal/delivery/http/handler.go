package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/bhankey/pharmacy-automatization/internal/delivery/http/v1/models"
	"github.com/bhankey/pharmacy-automatization/pkg/logger"
	"github.com/sirupsen/logrus"
)

type BaseHandler struct {
	Logger logger.Logger
}

func NewHandler(l logger.Logger) *BaseHandler {
	h := &BaseHandler{
		Logger: l,
	}

	return h
}

func (h *BaseHandler) WriteErrorResponse(ctx context.Context, w http.ResponseWriter, err error, isShown bool) {
	h.Logger.WithFields(logrus.Fields{
		"error":   err,
		"context": ctx,
	}).Errorf("response.error")

	w.WriteHeader(http.StatusBadRequest)

	var resp models.BaseResponse
	if err == nil || !isShown {
		resp = models.BaseResponse{
			Error:   "Something went wrong",
			Success: false,
		}
	} else {
		resp = models.BaseResponse{
			Error:   err.Error(),
			Success: false,
		}
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

func WriteResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

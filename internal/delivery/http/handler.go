package http

import (
	"github.com/bhankey/BD_lab/backend/pkg/logger"
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

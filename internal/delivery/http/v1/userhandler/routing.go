package userhandler

import (
	"context"

	deliveryhttp "github.com/bhankey/BD_lab/backend/internal/delivery/http"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Router  chi.Router
	userSrv UserSrv

	*deliveryhttp.BaseHandler
}

type UserSrv interface {
	Login(ctx context.Context, email string, password string) (string, error)
}

func NewUserHandler(baseHandler *deliveryhttp.BaseHandler, userSrv UserSrv) *UserHandler {
	router := chi.NewRouter()

	handeler := &UserHandler{
		Router:      router,
		userSrv:     userSrv,
		BaseHandler: baseHandler,
	}

	handeler.initRoutes(router)

	return handeler
}

func (h *UserHandler) initRoutes(router chi.Router) {
	router.Post("/login", h.login)
}

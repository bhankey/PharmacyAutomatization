package userhandler

import (
	"context"

	deliveryhttp "github.com/bhankey/pharmacy-automatization/internal/delivery/http"
	"github.com/bhankey/pharmacy-automatization/internal/entities"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	Router  chi.Router
	userSrv UserSrv

	*deliveryhttp.BaseHandler
}

type UserSrv interface {
	Login(ctx context.Context, email, password string, identifyData entities.UserIdentifyData) (entities.Tokens, error)
	RefreshToken(ctx context.Context, refreshToken string, identifyData entities.UserIdentifyData) (entities.Tokens, error)
	// ResetPassword(ctx context.Context, email, code, newPassword string) error
}

func NewUserHandler(baseHandler *deliveryhttp.BaseHandler, userSrv UserSrv) *UserHandler {
	router := chi.NewRouter()

	handler := &UserHandler{
		Router:      router,
		userSrv:     userSrv,
		BaseHandler: baseHandler,
	}

	handler.initRoutes(router)

	return handler
}

func (h *UserHandler) initRoutes(router chi.Router) {
	router.Post("/login", h.login)
	router.Post("/refresh", h.refresh)
	router.Get("/some", h.resetPassword)
}

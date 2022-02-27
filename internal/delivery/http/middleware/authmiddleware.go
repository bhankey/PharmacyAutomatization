package middleware

import (
	"context"
	"github.com/bhankey/BD_lab/backend/internal/entities"
	"github.com/bhankey/BD_lab/backend/pkg/logger"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type AuthMiddleware struct {
	jwtKey string

	logger logger.Logger
}

func NewAuthMiddleware(log logger.Logger, jwtKey string) *AuthMiddleware {
	return &AuthMiddleware{
		jwtKey: jwtKey,
		logger: log,
	}
}

func (m *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			if _, err := w.Write([]byte("failed to authorize user")); err != nil {
				m.logger.Warn("failed to write to user", err)
			}

			return
		}

		claim := entities.Claims{}
		token, err := jwt.ParseWithClaims(authHeader, &claim, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.jwtKey), nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			if _, err := w.Write([]byte("failed to authorize user")); err != nil {
				m.logger.Warn("failed to write to user", err)
			}

			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "client_id", claim.UserID)
		ctx = context.WithValue(ctx, "email", claim.Email)

		r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

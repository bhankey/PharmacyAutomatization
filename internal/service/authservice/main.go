package authservice

import (
	"context"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

type AuthService struct {
	userStorage  UserStorage
	tokenStorage TokenStorage

	jwtKey string
}

type UserStorage interface {
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	GetUserByID(ctx context.Context, id int) (entities.User, error)
}

type TokenStorage interface {
	CreateRefreshToken(ctx context.Context, token entities.RefreshToken) error
	GetAllActiveRefreshTokens(ctx context.Context, userID int) ([]entities.RefreshToken, error)
	DeactivateTokenByIDs(ctx context.Context, tokenIDs []int) error
	GetToken(ctx context.Context, refreshToken string) (entities.RefreshToken, error)
}

func NewAuthService(
	userStorage UserStorage,
	tokenStorage TokenStorage,
	jwtKey string,
) *AuthService {
	return &AuthService{
		userStorage:  userStorage,
		tokenStorage: tokenStorage,
		jwtKey:       jwtKey,
	}
}

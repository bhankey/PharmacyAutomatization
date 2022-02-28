package userservice

import (
	"context"
	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

type UserService struct {
	userStorage UserStorage
	tokenStorage

	passwordSalt string
	jwtKey       string
}

type UserStorage interface {
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
}

type tokenStorage interface {
	InsertRefreshToken(ctx context.Context, token entities.RefreshToken) error
	GetAllActiveRefreshTokens(ctx context.Context, userID int) ([]entities.RefreshToken, error)
	DeactivateTokenByIDs(ctx context.Context, tokenIDs []int) error
}

func NewUserService(userStorage UserStorage, passwordSalt string, jwtKey string) *UserService {
	return &UserService{
		userStorage:  userStorage,
		passwordSalt: passwordSalt,
		jwtKey:       jwtKey,
	}
}

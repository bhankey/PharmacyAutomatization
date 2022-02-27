package userservice

import (
	"context"
	"github.com/bhankey/BD_lab/backend/internal/entities"
)

type UserService struct {
	userStorage UserStorage

	passwordSalt string
	jwtKey       string
}

type UserStorage interface {
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	GetAllActiveRefreshTokens(ctx context.Context, userID int)
}

func NewUserService(userStorage UserStorage, passwordSalt string, jwtKey string) *UserService {
	return &UserService{
		userStorage:  userStorage,
		passwordSalt: passwordSalt,
		jwtKey:       jwtKey,
	}
}

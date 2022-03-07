package authservice

import (
	"context"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

type AuthService struct {
	userStorage          UserStorage
	tokenStorage         TokenStorage
	emailStorage         EmailStorage
	oneTimesCodesStorage OneTimesCodesStorage

	passwordSalt string
	jwtKey       string
}

type UserStorage interface {
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	GetUserByID(ctx context.Context, id int) (entities.User, error)
	CreateUser(ctx context.Context, user entities.User) error
	UpdatePassword(ctx context.Context, email string, newPasswordHash string) error
}

type TokenStorage interface {
	CreateRefreshToken(ctx context.Context, token entities.RefreshToken) error
	GetAllActiveRefreshTokens(ctx context.Context, userID int) ([]entities.RefreshToken, error)
	DeactivateTokenByIDs(ctx context.Context, tokenIDs []int) error
	GetToken(ctx context.Context, refreshToken string) (entities.RefreshToken, error)
}

type EmailStorage interface {
	SendResetPasswordCode(email string, code string) error
}

type OneTimesCodesStorage interface {
	CreateResetPasswordCode(ctx context.Context, email string, code string) error
	DeleteResetPasswordCode(ctx context.Context, email string) error
	GetResetPasswordCode(ctx context.Context, email string) (string, error)
}

func NewUserService(
	userStorage UserStorage,
	tokenStorage TokenStorage,
	emailStorage EmailStorage,
	oneTimesCodesStorage OneTimesCodesStorage,
	passwordSalt string,
	jwtKey string,
) *AuthService {
	return &AuthService{
		userStorage:          userStorage,
		tokenStorage:         tokenStorage,
		emailStorage:         emailStorage,
		oneTimesCodesStorage: oneTimesCodesStorage,
		passwordSalt:         passwordSalt,
		jwtKey:               jwtKey,
	}
}

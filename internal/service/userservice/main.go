package userservice

import (
	"context"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

type UserService struct {
	userStorage          UserStorage
	emailStorage         EmailStorage
	oneTimesCodesStorage OneTimesCodesStorage
}

type UserStorage interface {
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	CreateUser(ctx context.Context, user entities.User) error
	UpdatePassword(ctx context.Context, email string, newPasswordHash string) error
	GetBatchOfUsers(ctx context.Context, lastClientID int, limit int) ([]entities.User, error)
	UpdateUser(ctx context.Context, user entities.User) error
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
	emailStorage EmailStorage,
	oneTimesCodesStorage OneTimesCodesStorage,
) *UserService {
	return &UserService{
		userStorage:          userStorage,
		emailStorage:         emailStorage,
		oneTimesCodesStorage: oneTimesCodesStorage,
	}
}

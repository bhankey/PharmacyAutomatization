package userservice

import (
	"context"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) Registry(ctx context.Context, user entities.User) error {
	errBase := fmt.Sprintf("userservice.Registry(%v)", user)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s: failed to create hash from password error: %w", errBase, err)
	}

	user.PasswordHash = string(passwordHash)

	if err := s.userStorage.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("%s: failed to create user error: %w", errBase, err)
	}

	return nil
}

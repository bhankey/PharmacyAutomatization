package authservice

import (
	"context"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Registry(ctx context.Context, user entities.User, identifyData entities.UserIdentifyData) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to create hash from password error: %w", err)
	}

	user.PasswordHash = string(passwordHash)

	if err := s.userStorage.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("failed to create user error: %w", err)
	}

	return nil
}

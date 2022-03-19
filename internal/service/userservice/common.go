package userservice

import (
	"context"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (s *UserService) UpdateUser(ctx context.Context, user entities.User) error {
	errBase := fmt.Sprintf("userservice.UpdateUser(%v)", user)

	if err := s.userStorage.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("%s: failed to update user: %w", errBase, err)
	}

	return nil
}

func (s *UserService) GetBatchOfUsers(ctx context.Context, lastClientID int, limit int) ([]entities.User, error) {
	errBase := fmt.Sprintf("userservice.GetBatchOfUsers(%d, %d)", lastClientID, limit)

	users, err := s.userStorage.GetBatchOfUsers(ctx, lastClientID, limit)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to get users: %w", errBase, err)
	}

	return users, nil
}

package userservice

import (
	"context"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (s *UserService) UpdateUser(ctx context.Context, user entities.User) error {
	return s.userStorage.UpdateUser(ctx, user)
}

func (s *UserService) GetBatchOfUsers(ctx context.Context, lastClientID int, limit int) ([]entities.User, error) {
	return s.userStorage.GetBatchOfUsers(ctx, lastClientID, limit)
}

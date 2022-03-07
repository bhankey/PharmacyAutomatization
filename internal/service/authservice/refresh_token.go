package authservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/apperror"
	"github.com/bhankey/pharmacy-automatization/internal/entities"
)

func (s *AuthService) RefreshToken(
	ctx context.Context,
	refreshToken string,
	identifyData entities.UserIdentifyData,
) (entities.Tokens, error) {
	token, err := s.tokenStorage.GetToken(ctx, refreshToken)
	if err != nil {
		if errors.Is(err, apperror.ErrNoEntity) {
			return entities.Tokens{}, apperror.NewClientError(apperror.WrongAuthToken, err)
		}

		return entities.Tokens{}, fmt.Errorf("failed to get refresh token error: %w", err)
	}

	user, err := s.userStorage.GetUserByID(ctx, token.UserID)
	if err != nil {
		return entities.Tokens{}, fmt.Errorf("failed to get user error: %w", err)
	}

	accessToken, err := s.createAndSignedToken(user.ID, user.Email, jwtExpireTime)
	if err != nil {
		return entities.Tokens{}, fmt.Errorf("failed to create access token error: %w", err)
	}

	newRefreshToken, err := s.createAndSaveRefreshToken(ctx, user.ID, user.Email, identifyData)
	if err != nil {
		return entities.Tokens{}, fmt.Errorf("failed to create refresh token error: %w", err)
	}

	// TODO write wrap to catch panic in gorutine
	go func() {
		_ = s.tokenStorage.DeactivateTokenByIDs(ctx, []int{token.ID}) // TODO think about async error handling
	}()

	return entities.Tokens{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

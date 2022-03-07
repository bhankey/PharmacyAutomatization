package authservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/apperror"
	"github.com/bhankey/pharmacy-automatization/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

const lengthOfCode = 6

func (s *AuthService) ResetPassword(ctx context.Context, email, code, newPassword string) error {
	cachedCode, err := s.oneTimesCodesStorage.GetResetPasswordCode(ctx, email)
	if err != nil {
		if errors.Is(err, apperror.ErrNoEntity) {
			return apperror.NewClientError(apperror.WrongOneTimeCode, err)
		}

		return err
	}

	if cachedCode != code {
		return apperror.NewClientError(apperror.WrongOneTimeCode, err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := s.userStorage.UpdatePassword(ctx, email, string(passwordHash)); err != nil {
		return err
	}

	// TODO write wrap to catch panic in gorutine
	go func() {
		_ = s.oneTimesCodesStorage.DeleteResetPasswordCode(ctx, email) // TODO think about async error handling
	}()

	return nil
}

func (s *AuthService) RequestToResetPassword(ctx context.Context, email string) error {
	_, err := s.userStorage.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, apperror.ErrNoEntity) {
			return apperror.NewClientError(apperror.NoClient, err)
		}

		return fmt.Errorf("no client for this email: %w", err)
	}

	randomCode := utils.RandStringBytes(lengthOfCode)

	if err := s.oneTimesCodesStorage.CreateResetPasswordCode(ctx, email, randomCode); err != nil {
		return err
	}

	if err := s.emailStorage.SendResetPasswordCode(email, randomCode); err != nil {
		return err
	}

	return nil
}

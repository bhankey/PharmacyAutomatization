package userservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/apperror"
	"github.com/bhankey/pharmacy-automatization/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

const lengthOfCode = 6

func (s *UserService) ResetPassword(ctx context.Context, email, code, newPassword string) error {
	errBase := fmt.Sprintf("userservice.ResetPassword(%s, %s, %s)", email, code, newPassword)

	cachedCode, err := s.oneTimesCodesStorage.GetResetPasswordCode(ctx, email)
	if err != nil {
		if errors.Is(err, apperror.ErrNoEntity) {
			return apperror.NewClientError(apperror.WrongOneTimeCode, err)
		}

		return fmt.Errorf("%s: failed to get reset code: %w", errBase, err)
	}

	if cachedCode != code {
		return apperror.NewClientError(apperror.WrongOneTimeCode, err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%s: failed to generete hash: %w", errBase, err)
	}

	if err := s.userStorage.UpdatePassword(ctx, email, string(passwordHash)); err != nil {
		return fmt.Errorf("%s failed to update password: %w", errBase, err)
	}

	go func() {
		ctx := context.Background()

		_ = s.oneTimesCodesStorage.DeleteResetPasswordCode(ctx, email)
	}()

	return nil
}

func (s *UserService) RequestToResetPassword(ctx context.Context, email string) error {
	errBase := fmt.Sprintf("userservice.RequestToResetPassword(%s)", email)

	_, err := s.userStorage.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, apperror.ErrNoEntity) {
			return apperror.NewClientError(apperror.NoClient, err)
		}

		return fmt.Errorf("%s: no client for this email: %w", errBase, err)
	}

	randomCode := utils.RandStringBytes(lengthOfCode)

	if err := s.oneTimesCodesStorage.CreateResetPasswordCode(ctx, email, randomCode); err != nil {
		return fmt.Errorf("%s: failed to create reset code: %w", errBase, err)
	}

	if err := s.emailStorage.SendResetPasswordCode(email, randomCode); err != nil {
		return fmt.Errorf("%s: fialed to send password code email: %w", errBase, err)
	}

	return nil
}

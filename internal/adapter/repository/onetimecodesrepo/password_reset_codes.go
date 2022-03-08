package onetimecodesrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/bhankey/pharmacy-automatization/internal/apperror"
	"github.com/go-redis/redis/v8"
)

func (r *ResetCodesRepo) CreateResetPasswordCode(ctx context.Context, email string, code string) error {
	errBase := fmt.Sprintf("onetimecodesrepo.CreateResetPasswordCode(%s, %s)", email, code)

	if err := r.redis.Set(ctx, "reset_password_"+email, code, r.resetPasswordTTL).Err(); err != nil {
		return fmt.Errorf("%s: failed to create reset code: %w", errBase, err)
	}

	return nil
}

func (r *ResetCodesRepo) DeleteResetPasswordCode(ctx context.Context, email string) error {
	errBase := fmt.Sprintf("onetimecodesrepo.DeleteResetPasswordCode(%s)", email)

	if err := r.redis.Del(ctx, "reset_password_"+email).Err(); err != nil {
		return fmt.Errorf("%s :failed to delete reset password code: %w", errBase, err)
	}

	return nil
}

func (r *ResetCodesRepo) GetResetPasswordCode(ctx context.Context, email string) (string, error) {
	errBase := fmt.Sprintf("onetimecodesrepo.GetResetPasswordCode(%s)", email)

	code, err := r.redis.Get(ctx, "reset_password_"+email).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", apperror.ErrNoEntity
		}

		return "", fmt.Errorf("%s: failed to get reset password code: %w", errBase, err)
	}

	return code, nil
}

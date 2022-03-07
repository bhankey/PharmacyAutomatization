package onetimecodesrepo

import (
	"context"
	"errors"

	"github.com/bhankey/pharmacy-automatization/internal/apperror"
	"github.com/go-redis/redis/v8"
)

func (r *ResetCodesRepo) CreateResetPasswordCode(ctx context.Context, email string, code string) error {
	if err := r.redis.Set(ctx, "reset_password_"+email, code, r.resetPasswordTTL).Err(); err != nil {
		return err
	}

	return nil
}

func (r *ResetCodesRepo) DeleteResetPasswordCode(ctx context.Context, email string) error {
	if err := r.redis.Del(ctx, "reset_password_"+email).Err(); err != nil {
		return err
	}

	return nil
}

func (r *ResetCodesRepo) GetResetPasswordCode(ctx context.Context, email string) (string, error) {
	code, err := r.redis.Get(ctx, "reset_password_"+email).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", apperror.ErrNoEntity
		}

		return "", err
	}

	return code, err
}

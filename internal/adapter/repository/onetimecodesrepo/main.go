package onetimecodesrepo

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type ResetCodesRepo struct {
	resetPasswordTTL time.Duration

	redis *redis.Client
}

func NewResetCodesRepo(redis *redis.Client, resetPasswordTTL time.Duration) *ResetCodesRepo {
	return &ResetCodesRepo{
		resetPasswordTTL: resetPasswordTTL,
		redis:            redis,
	}
}

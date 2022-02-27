package userservice

import (
	"context"
	"fmt"
	"github.com/bhankey/BD_lab/backend/internal/entities"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const jwtExpireTime = time.Minute * 15
const jwtExpireRefreshTime = time.Hour * 24 * 60

func (s *UserService) Login(ctx context.Context, email string, password string) (string, error) {
	user, err := s.userStorage.GetUserByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("failed to get password from hash error: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", fmt.Errorf("wrong password error: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &entities.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtExpireTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Email:  email,
		UserID: user.ID,
	})

	return token.SignedString([]byte(s.jwtKey))
}

func (s *UserService) GenerateAndSaveRefreshToken(ctx context.Context, userID int, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &entities.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtExpireRefreshTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Email:  email,
		UserID: userID,
	})

	signedToken, err := token.SignedString([]byte(s.jwtKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token error: %w", err)
	}

}

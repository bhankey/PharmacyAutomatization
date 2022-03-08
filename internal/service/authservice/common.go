package authservice

import (
	"context"
	"time"

	"github.com/bhankey/pharmacy-automatization/internal/entities"
	"github.com/dgrijalva/jwt-go"
)

func (s *AuthService) createAndSaveRefreshToken(
	ctx context.Context,
	userID int,
	email string,
	identifyData entities.UserIdentifyData,
) (string, error) {
	signedToken, err := s.createAndSignedToken(userID, email, jwtExpireRefreshTime)
	if err != nil {
		return "", err
	}

	refreshToken := entities.RefreshToken{
		UserID:      userID,
		Token:       signedToken,
		UserAgent:   identifyData.UserAgent,
		IP:          identifyData.IP,
		FingerPrint: identifyData.FingerPrint,
	}

	if err := s.tokenStorage.CreateRefreshToken(ctx, refreshToken); err != nil {
		return "", err // nolint: wrapcheck, nolintlint
	}

	return signedToken, nil
}

func (s *AuthService) createAndSignedToken(userID int, email string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &entities.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Email:  email,
		UserID: userID,
	})

	signedToken, err := token.SignedString([]byte(s.jwtKey))
	if err != nil {
		return "", err // nolint: wrapcheck, nolintlint
	}

	return signedToken, nil
}

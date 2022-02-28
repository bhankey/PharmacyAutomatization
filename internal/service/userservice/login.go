package userservice

import (
	"context"
	"fmt"
	"github.com/bhankey/pharmacy-automatization/internal/entities"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"sort"
	"time"
)

const jwtExpireTime = time.Minute * 15
const jwtExpireRefreshTime = time.Hour * 24 * 60
const maxActivateToken = 5

// Move this to some service that implement registry pattern (consul)

func (s *UserService) Login(ctx context.Context, email string, password string) (entities.Tokens, error) {
	user, err := s.userStorage.GetUserByEmail(ctx, email)
	if err != nil {
		return entities.Tokens{}, fmt.Errorf("failed to get password from hash error: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return entities.Tokens{}, fmt.Errorf("wrong password error: %w", err)
	}

	accessToken, err := s.createAndSignedToken(user.ID, email, jwtExpireTime)
	if err != nil {
		return entities.Tokens{}, fmt.Errorf("failed to create and sigend accass token error: %w", err)
	}

	refreshToken, err := s.GenerateAndSaveRefreshToken(ctx, user.ID, email) // TODO
	if err != nil {
		return entities.Tokens{}, fmt.Errorf("failed to create and sigend refresh token error: %w", err)
	}

	return entities.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken, // TODO
	}, nil
}

func (s *UserService) GenerateAndSaveRefreshToken(ctx context.Context, userID int, email, userAgent, fingerPrint, ip string) (string, error) {
	errorBase := fmt.Sprintf("userservice.GenerateAndSaveRefreshToken(user_id: %d, email: %s, user_agent: %s, finger_print: %s, ip: %s)", userID, email, userAgent, fingerPrint, ip)

	signedToken, err := s.createAndSignedToken(userID, email, jwtExpireRefreshTime)
	if err != nil {
		return "", fmt.Errorf("%s.createAndSignedToken.error: %w", errorBase, err)
	}

	refreshToken := entities.RefreshToken{
		UserID:      userID,
		Token:       signedToken,
		UserAgent:   userAgent,
		IP:          ip,
		FingerPrint: fingerPrint,
	}

	if err := s.tokenStorage.InsertRefreshToken(ctx, refreshToken); err != nil {
		return "", fmt.Errorf("%s.InsertRefreshToken.error: %w", errorBase, err)
	}

	// TODO write wrap to catch panic in gorutine
	go func() {
		time.Sleep(time.Second * 20) // In 20 seconds slave will surely update data

		_ = s.deactivateMaxReachedTokensCount(ctx, userID) // TODO think about async error handling
	}()

	return signedToken, nil
}

func (s *UserService) deactivateMaxReachedTokensCount(ctx context.Context, userID int) error {
	tokens, err := s.tokenStorage.GetAllActiveRefreshTokens(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to get all active tokens error: %w", err)
	}

	if len(tokens) > maxActivateToken {
		tokensIDsToDeativate := make([]int, 0, len(tokens)-maxActivateToken)
		sort.Slice(tokens, func(i, j int) bool {
			return tokens[i].ID > tokens[j].ID
		})

		for i := 0; i < len(tokens)-maxActivateToken; i++ {
			tokensIDsToDeativate = append(tokensIDsToDeativate, tokens[i].ID)
		}

		if err := s.tokenStorage.DeactivateTokenByIDs(ctx, tokensIDsToDeativate); err != nil {
			return fmt.Errorf("failed to deativate old tokens error: %w", err)
		}
	}

	return nil
}

func (s *UserService) createAndSignedToken(userID int, email string, ttl time.Duration) (string, error) {
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
		return "", fmt.Errorf("failed to sign token error: %w", err)
	}

	return signedToken, nil
}

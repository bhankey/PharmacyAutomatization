package entities

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	Email  string `json:"email"`
	UserID int    `json:"user_id"`
	Role   Role   `json:"role"`
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type RefreshToken struct {
	ID           int
	UserID       int
	Token        string
	UserAgent    string
	IP           string
	FingerPrint  string
	IsAvailable  bool
	CreationTime time.Time
}

package entities

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	Email  string `json:"email"`
	UserID int    `json:"user_id"`
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

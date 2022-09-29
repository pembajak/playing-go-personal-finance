package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Payload define payload body for token
type Payload struct {
	ID       int64
	FullName string
	Phone    string
	Email    string
}

type Claims struct {
	Payload
	jwt.RegisteredClaims
}

type GetToken struct {
	RefreshToken          string
	AccessToken           string
	RefreshTokenExpiresAt time.Time
	AccessTokenExpiresAt  time.Time
}

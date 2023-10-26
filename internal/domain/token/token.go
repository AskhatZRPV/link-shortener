package token

import (
	"link-shortener/internal/domain/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	SessID    string
	UserID    user.ID
	ExpiresAt time.Time
}

func NewClaimsFactory(sessID string, userID user.ID) func(expiresAt time.Time) *Claims {
	return func(expiresAt time.Time) *Claims {
		return &Claims{sessID, userID, expiresAt}
	}
}

type Token string

func New(token string) *Token {
	var f Token = Token(token)
	return &f
}

type JwtAccessClaims struct {
	ID        string `json:"id"`
	ExpiresAt int64  `json:"exp"`
	jwt.MapClaims
}

type JwtRefreshClaims struct {
	ID        string `json:"id"`
	ExpiresAt int64  `json:"exp"`
	jwt.MapClaims
}

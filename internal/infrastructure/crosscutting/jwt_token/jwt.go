package jwt_token

import (
	"context"
	"link-shortener/internal/core/config"
	"link-shortener/internal/domain/token"
	"link-shortener/internal/domain/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type provider struct {
	authCfg config.Auth
}

func New(config *config.Config) token.Provider {
	return &provider{config.Auth}
}

func (p *provider) Grant(ctx context.Context, userID user.ID) (r *token.GrantResult, err error) {
	at, err := p.createAccessToken(string(userID), p.authCfg.AccessTokenSecret)
	if err != nil {
		return nil, err
	}
	rt, err := p.createRefreshToken(string(userID), p.authCfg.RefreshTokenSecret)
	if err != nil {
		return nil, err
	}
	return &token.GrantResult{Access: at, Refresh: rt}, nil
}

func (p *provider) createAccessToken(userID, secret string) (string, error) {
	claimsRefresh := &token.JwtAccessClaims{
		ID:        userID,
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, err

}

func (p *provider) createRefreshToken(userID, secret string) (string, error) {
	claimsRefresh := &token.JwtRefreshClaims{
		ID:        userID,
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, err

}

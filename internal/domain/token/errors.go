package token

import "link-shortener/internal/core/domainerr"

var (
	ErrTokenExpired = domainerr.New("token expired")
	ErrInvalidSign  = domainerr.New("invalid sign")
)

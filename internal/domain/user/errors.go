package user

import "link-shortener/internal/core/domainerr"

var (
	ErrUserExitst   = domainerr.New("user exists")
	ErrUserNotFound = domainerr.New("user not found")
)

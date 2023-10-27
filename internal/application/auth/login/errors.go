package login

import "link-shortener/internal/core/domainerr"

var (
	ErrAccountDoesNotExist = domainerr.New("account does not exist")
	ErrIncorrectPassword   = domainerr.New("err incorrect password")
)

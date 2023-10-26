package register

import "link-shortener/internal/core/domainerr"

var (
	ErrAccountAlreadyExists = domainerr.New("account already exists")
)

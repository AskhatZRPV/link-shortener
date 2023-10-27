package find

import "link-shortener/internal/core/domainerr"

var (
	ErrLinkDoesNotExist  = domainerr.New("account does not exist")
	ErrLinkAlreadyExists = domainerr.New("link already exist")
)

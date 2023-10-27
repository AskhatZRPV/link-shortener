package link

import "link-shortener/internal/core/domainerr"

var (
	ErrLinkExitst   = domainerr.New("link exists")
	ErrLinkNotFound = domainerr.New("link not found")
)

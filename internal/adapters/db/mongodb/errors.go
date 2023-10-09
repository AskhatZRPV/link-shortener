package mongodb

import "errors"

var (
	// ErrNotAuthorized is used when the the access is not permisted
	ErrNotAuthorized = errors.New("not authorized")
	// ErrLinkNotFound is used when an link is not found
	ErrLinkNotFound = errors.New("link not found")
	// ErrLinkExists is used when an link already exists
	ErrLinkExists = errors.New("link already exists")
	// ErrInvalidAdminRouter is used when an invalid admin router is given
	ErrInvalidAdminRouter = errors.New("invalid admin router given")
)

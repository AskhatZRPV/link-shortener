package mongodb

import "errors"

var (
	// ErrNotAuthorized is used when the the access is not permisted
	ErrNotAuthorized = errors.New("not authorized")
	// ErrUserNotFound is used when an user is not found
	ErrLinkNotFound = errors.New("link not found")
	// ErrUserExists is used when an user already exists
	ErrLinkExists = errors.New("link already exists")
	// ErrInvalidAdminRouter is used when an invalid admin router is given
	ErrInvalidAdminRouter = errors.New("invalid admin router given")
)

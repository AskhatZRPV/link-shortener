package usermongo

import "errors"

var (
	// ErrNotAuthorized is used when the the access is not permisted
	ErrNotAuthorized = errors.New("not authorized")
	// ErrUserNotFound is used when an user is not found
	ErrUserNotFound = errors.New("user not found")
	// ErrUserExists is used when an user already exists
	ErrUserExists = errors.New("user already exists")
	// ErrInvalidAdminRouter is used when an invalid admin router is given
	ErrInvalidAdminRouter = errors.New("invalid admin router given")
)

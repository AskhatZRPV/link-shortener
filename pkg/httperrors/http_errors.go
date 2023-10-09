package httperrors

import (
	"errors"
	"time"
)

var (
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrInvalidPayload = errors.New("invalid payload")
)

type HttpError struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func BuildHttpError(err error) HttpError {
	return HttpError{
		Message:   err.Error(),
		Timestamp: time.Now(),
	}
}

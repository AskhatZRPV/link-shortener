package link

import (
	"time"

	"github.com/google/uuid"
)

type ID string

type Link struct {
	ID
	Alias       string
	RedirectUrl string
	// UserID      user.ID
	CreatedAt *time.Time
}

func New(alias, redirectUrl string) *Link {
	now := time.Now()
	return &Link{
		ID:          ID(uuid.New().String()),
		Alias:       alias,
		RedirectUrl: redirectUrl,
		CreatedAt:   &now,
	}
}

func FromData(id ID, alias string, redirectUrl string, created *time.Time) *Link {
	return &Link{
		id,
		alias,
		redirectUrl,
		created,
	}
}

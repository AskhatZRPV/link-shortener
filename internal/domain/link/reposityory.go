package link

import (
	"context"
)

type Result struct {
	ID          string
	Alias       string
	Domain      string
	RedirectUrl string
}

type Repository interface {
	Save(context.Context, *Link) error
	FindByAlias(ctx context.Context, alias string) (*Link, error)
}

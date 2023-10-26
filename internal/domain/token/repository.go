package token

import (
	"context"
	"link-shortener/internal/domain/user"
)

type GrantResult struct {
	Access, Refresh string
}

type Provider interface {
	Grant(context.Context, user.ID) (*GrantResult, error)
}

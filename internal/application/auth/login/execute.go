package login

import (
	"context"
	"link-shortener/internal/core/domainerr"
	"link-shortener/internal/domain/session"
	"link-shortener/internal/domain/token"
	"link-shortener/internal/domain/user"
	"time"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*Result, error) {
	u, g, err := i.loginAndGrantTokens(ctx, p)
	if err != nil {
		return nil, err
	}

	err = i.txManager.Do(ctx, func(ctx context.Context) error {
		err := i.refreshRepo.Save(ctx, session.New(g.Refresh, u.ID, time.Now().Add(15*time.Minute)))
		if err != nil {
			return errors.Wrap(err, "failed to persist refresh token")
		}

		err = i.userRepo.UpdateLastLoginAttempt(ctx, u.ID, time.Now())
		if err != nil {
			return errors.Wrap(err, "faied to update last login attempt")
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to run transaction")
	}

	return g, nil
}

func (i *implementation) loginAndGrantTokens(ctx context.Context, p *Payload) (*user.User, *token.GrantResult, error) {
	u, err := i.userRepo.FindByUsername(ctx, p.Username)
	if err != nil {
		return nil, nil, domainerr.Join(ErrAccountDoesNotExist, err)
	}

	if correctPassword := i.ph.Compare(p.Password, u.Password); !correctPassword {
		return nil, nil, errors.Wrap(ErrIncorrectPassword, "incorrect password received")
	}

	g, err := i.tokenProvider.Grant(ctx, u.ID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to grant tokens")
	}

	return u, g, nil
}

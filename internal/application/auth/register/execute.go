package register

import (
	"context"
	"link-shortener/internal/domain/user"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) error {
	_, err := i.userRepo.FindByUsername(ctx, p.Username)
	if err == nil || !errors.Is(err, user.ErrUserNotFound) {
		return errors.Wrap(ErrAccountAlreadyExists, "account with such username exists")
	}
	pwd, err := i.ph.Hash(p.Password)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}

	err = i.userRepo.Save(ctx, user.New(p.Username, pwd))
	if err != nil {
		return errors.Wrap(err, "failed to save user in repository")
	}

	return nil
}

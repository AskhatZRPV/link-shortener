package create

import (
	"context"
	"link-shortener/internal/domain/link"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) error {
	_, err := i.linkRepo.FindByAlias(ctx, p.Alias)
	if nil == err || !errors.Is(err, link.ErrLinkNotFound) {
		return errors.Wrap(err, "account with alias exists")
	}

	err = i.linkRepo.Save(ctx, link.New(p.Alias, p.RedirectUrl))
	if err != nil {
		return errors.Wrap(err, "failed to save user in repository")
	}

	return nil
}

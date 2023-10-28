package create

import (
	"context"
	"link-shortener/internal/domain/link"
	"link-shortener/pkg/utils/randomalias"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) error {
	var alias string
	if len(p.Alias) != 0 {
		_, err := i.linkRepo.FindByAlias(ctx, p.Alias)
		if err == nil || !errors.Is(err, link.ErrLinkNotFound) {
			return errors.Wrap(err, "link with alias exists")
		}
		alias = p.Alias
	} else {
		for {
			alias = randomalias.NewRandomString(6)
			_, err := i.linkRepo.FindByAlias(ctx, alias)
			if err != nil {
				break
			}
		}
	}

	err := i.linkRepo.Save(ctx, link.New(alias, p.RedirectUrl))
	if err != nil {
		return errors.Wrap(err, "failed to save link in repository")
	}

	return nil
}

package find

import (
	"context"
	"link-shortener/internal/domain/link"

	"github.com/pkg/errors"
)

func (i *implementation) Execute(ctx context.Context, p *Payload) (*link.Link, error) {
	res, err := i.linkRepo.FindByAlias(ctx, p.Alias)
	if err != nil {
		return nil, errors.Wrap(err, "link not found")
	}
	return res, nil
}

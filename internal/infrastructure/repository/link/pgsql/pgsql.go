package pgsql

import (
	"context"
	"database/sql"
	"link-shortener/internal/domain/link"
	"link-shortener/internal/infrastructure/repository/tx/pgsqltx"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) link.Repository {
	return &repo{db}
}

func (i *repo) Save(ctx context.Context, l *link.Link) error {
	const insertLinkQuery = `
		INSERT INTO links (id, alias, redirect_url) VALUES($1, $2, $3);
	`

	q := pgsqltx.QuerierFromCtx(ctx, i.db)
	if _, err := q.ExecContext(ctx, insertLinkQuery, l.ID, l.Alias, l.RedirectUrl); err != nil {
		return errors.Wrap(err, "failed to insert new link record")
	}

	return nil
}

func (i *repo) FindByAlias(ctx context.Context, alias string) (*link.Link, error) {
	const selectLinkByAliasAndDomainQuery = `
		SELECT * FROM links WHERE alias LIKE $1;
	`

	q := pgsqltx.QuerierFromCtx(ctx, i.db)
	var row linkRow
	err := q.GetContext(ctx, &row, selectLinkByAliasAndDomainQuery, alias)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrap(link.ErrLinkNotFound, "link not found in pg repo")
		default:
			return nil, errors.Wrap(err, "unexpected query error")
		}
	}

	return row.ToDomain(), nil
}

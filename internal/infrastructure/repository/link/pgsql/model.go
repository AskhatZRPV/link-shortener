package pgsql

import (
	"link-shortener/internal/domain/link"
	"time"
)

type linkRow struct {
	ID          string     `db:"id"`
	Alias       string     `db:"alias"`
	RedirectUrl string     `db:"redirect_url"`
	CreatedAt   *time.Time `db:"created_at"`
}

func (r *linkRow) ToDomain() *link.Link {
	return link.FromData(link.ID(r.ID), r.Alias, r.RedirectUrl, r.CreatedAt)
}

package pgsql

import (
	"context"
	"link-shortener/internal/domain/session"
	"link-shortener/internal/infrastructure/repository/tx/pgsqltx"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) session.Repository {
	return &repo{db}
}

func (r *repo) Save(ctx context.Context, s *session.Session) error {
	const insertSessionQuery = `
		INSERT INTO user_sessions (id, user_id, expires_at)
		VALUES($1, $2, $3);
	`

	q := pgsqltx.QuerierFromCtx(ctx, r.db)
	_, err := q.ExecContext(ctx, insertSessionQuery, s.ID, s.UserID, s.ExpiresAt)
	if err != nil {
		return errors.Wrap(err, "failed to create session record")
	}

	return nil
}

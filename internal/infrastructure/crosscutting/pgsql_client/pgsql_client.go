package pgsql_client

import (
	"context"
	"link-shortener/internal/core/config"
	"time"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Querier interface {
	sqlx.QueryerContext
	sqlx.ExecerContext
	sqlx.ExtContext

	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

const connTimeout = time.Second * 3

func New(config *config.Config) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connTimeout)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "postgres", config.PGConnectionString)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to postgres")
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "failed to ping postgres")
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	return db, nil
}

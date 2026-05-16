package pgsql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Close()
	Pool() *pgxpool.Pool
	Ping(ctx context.Context) error
}

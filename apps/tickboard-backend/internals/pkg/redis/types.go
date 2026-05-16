package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Client interface {
	Close() error
	Rds() redis.UniversalClient
	Ping(ctx context.Context) error

	Count(ctx context.Context, pattern string) (int64, error)
	Exists(ctx context.Context, keyPrefix string) (bool, error)
	Destroy(ctx context.Context, pattern string) error
}

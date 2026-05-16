package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type client struct {
	redis  redis.UniversalClient
	config *Config
}

func NewClient(cfg *Config) (Client, error) {
	if cfg == nil {
		return nil, errors.New("redis: config is nil")
	}

	cfg.applyDefaults()
	opt := cfg.redisOption()

	rdb := redis.NewUniversalClient(opt)

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		return nil, fmt.Errorf("redis: unable to ping server: %w", err)
	}

	return &client{
		redis:  rdb,
		config: cfg,
	}, nil
}

func (c *client) Rds() redis.UniversalClient {
	return c.redis
}

func (c *client) Close() error {
	return c.redis.Close()
}

func (c *client) Ping(ctx context.Context) error {
	return c.redis.Ping(ctx).Err()
}

func (c *client) Count(ctx context.Context, pattern string) (int64, error) {
	iter := c.redis.Scan(ctx, 0, pattern, 100).Iterator()

	var count int64
	for iter.Next(ctx) {
		count++
	}

	return count, nil
}

func (c *client) Exists(ctx context.Context, keyPrefix string) (bool, error) {
	res, err := c.redis.Exists(ctx, keyPrefix).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return false, fmt.Errorf("redis: %w", err)
	}

	return res > 0, err
}

func (c *client) Destroy(ctx context.Context, pattern string) error {
	cursor := uint64(0)
	for {
		var keys []string
		var err error

		keys, cursor, err = c.redis.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			return fmt.Errorf("redis: %w", err)
		}

		if len(keys) == 0 {
			break
		}

		if err := c.redis.Del(ctx, keys...).Err(); err != nil {
			return fmt.Errorf("redis: %w", err)
		}
	}

	return nil
}

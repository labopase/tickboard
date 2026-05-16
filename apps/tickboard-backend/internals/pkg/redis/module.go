package redis

import (
	"context"
	"strconv"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/logger"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/utilities/uconfig"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

var Module = fx.Module("redis",
	fx.Provide(
		func() (*Config, error) {
			return uconfig.BindJSONKey[*Config]("redis")
		},
		NewClient,
		func(c Client) redis.UniversalClient {
			return c.Rds()
		},
	),
	fx.Invoke(
		func(lc fx.Lifecycle, c Client, log logger.Logger, cfg *Config) {
			lc.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						log.Infow("Redis connected",
							logger.Any("addrs", cfg.Addrs),
							logger.String("pool_size", strconv.Itoa(cfg.PoolSize)),
							logger.String("min_idle_conns", strconv.Itoa(cfg.MinIdleConns)),
							logger.String("dial_timeout", cfg.DialTimeout.String()),
							logger.String("read_timeout", cfg.ReadTimeout.String()),
							logger.String("write_timeout", cfg.WriteTimeout.String()),
						)
						return nil
					},
					OnStop: func(ctx context.Context) error {
						log.Info("Redis disconnected")
						return c.Close()
					},
				},
			)
		},
	),
)

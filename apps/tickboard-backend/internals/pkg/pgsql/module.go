package pgsql

import (
	"context"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/logger"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/utilities/uconfig"
	"go.uber.org/fx"
)

var Module = fx.Module("postgres",
	fx.Provide(
		func() (*Config, error) {
			return uconfig.BindJSONKey[*Config]("postgres")
		},
		NewClient,
	),
	fx.Invoke(
		func(lc fx.Lifecycle, c Client, log logger.Logger, cfg *Config) {
			lc.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						log.Infow("PostgreSQL connected", logger.String("host", cfg.Host), logger.Int("port", cfg.Port))
						return nil
					},
					OnStop: func(ctx context.Context) error {
						c.Close()
						log.Info("PostgreSQL disconnected")
						return nil
					},
				},
			)
		},
	),
)

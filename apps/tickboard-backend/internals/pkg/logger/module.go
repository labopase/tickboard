package logger

import (
	"context"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/config"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/utilities/uconfig"
	"go.uber.org/fx"
)

var Module = fx.Module("logger",
	fx.Provide(
		provideLoggerConfig,
		NewLogger,
	),
	fx.Invoke(
		func(lc fx.Lifecycle, log Logger, cfg *Config) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					log.Infow("Logger initialized",
						String("environment", cfg.Environment),
						Bool("enable_caller", cfg.EnableCaller),
						Bool("enable_trace", cfg.EnableTrace),
					)
					return nil
				},
				OnStop: func(ctx context.Context) error {
					log.Info("Logger shutting down")
					_ = log.Sync()
					return nil
				},
			})
		},
	),
)

func provideLoggerConfig(appCfg *config.Config) (*Config, error) {
	cfg, err := uconfig.BindJSONKey[*Config]("logger")
	if err != nil {
		return nil, err
	}

	cfg.Environment = appCfg.Environment

	return cfg, nil
}

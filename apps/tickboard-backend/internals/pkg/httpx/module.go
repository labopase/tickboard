package httpx

import (
	"context"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/logger"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/utilities/uconfig"
	"go.uber.org/fx"
)

var Module = fx.Module("http_engine",
	fx.Provide(
		provideHttpConfig,
		NewEngine,
	),
	fx.Invoke(registerHooks),
)

func provideHttpConfig() (*Config, error) {
	return uconfig.BindJSONKey[*Config]("server")
}

func registerHooks(lc fx.Lifecycle, e Engine, log logger.Logger, cfg *Config) {
	var (
		serverCancel context.CancelFunc
		serverDone   = make(chan error, 1)
	)

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Infow("initializing http server",
					logger.String("addr", cfg.Addr()),
					logger.String("read_timeout", cfg.ReadTimeout.String()),
					logger.String("idle_timeout", cfg.IdleTimeout.String()),
					logger.String("shutdown_timeout", cfg.ShutdownTimeout.String()),
				)

				serverCtx, cancel := context.WithCancel(context.Background())
				serverCancel = cancel

				go func() {
					serverDone <- e.Start(serverCtx)
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Infow("shutting down http server",
					logger.String("addr", cfg.Addr()),
				)

				serverCancel()

				select {
				case err := <-serverDone:
					return err
				case <-ctx.Done():
					return ctx.Err()
				}
			},
		},
	)
}

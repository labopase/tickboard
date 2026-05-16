package healthz

import (
	"context"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/httpx"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/pgsql"
	rds "github.com/halimdotnet/tickboard-backend/internals/pkg/redis"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/utilities/uconfig"
	"go.uber.org/fx"
)

var Module = fx.Module("health_checker",
	fx.Provide(
		func() (*Config, error) {
			return uconfig.BindJSONKey[*Config]("health")
		},
		NewChecker,
		NewHandler,
	),
	fx.Invoke(
		func(
			lc fx.Lifecycle,
			ck Checker,
			rd rds.Client,
			h Handler,
			pg pgsql.Client,
			cfg *Config,
			e httpx.Engine,
		) {
			if cfg.Redis {
				ck.Register("redis", RedisCheck(rd))
			}
			if cfg.Postgres {
				ck.Register("postgres", PostgresCheck(pg))
			}

			route := e.Instance().Group("/health")
			route.GET("/live", h.LivenessHandler)
			route.GET("/ready", h.ReadinessHandler)
			route.GET("/startup", h.StartupHandler)

			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					h.SetReady(true)
					return nil
				},
				OnStop: func(ctx context.Context) error {
					h.SetReady(false)
					return nil
				},
			})
		},
	),
)

package config

import (
	"errors"
	"slices"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/constants"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/utilities/uconfig"
	"go.uber.org/fx"
)

var Module = fx.Module("app_config",
	fx.Provide(
		func() (*Config, error) {
			cfg, err := uconfig.BindJSONKey[*Config]("app")
			if err != nil {
				return nil, err
			}

			if cfg.Name == "" {
				return nil, errors.New("app.name is required")
			}

			if cfg.Version == "" {
				return nil, errors.New("app.version is required")
			}

			if !slices.Contains(constants.ListEnv, cfg.Environment) {
				return nil, errors.New("app.environment is invalid")
			}

			return cfg, nil
		},
	),
)

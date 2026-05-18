package tickboard

import (
	"github.com/halimdotnet/tickboard-backend/internals/pkg/application"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/config"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/httpx"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/logger"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/pgsql"
	rds "github.com/halimdotnet/tickboard-backend/internals/pkg/redis"
	"github.com/halimdotnet/tickboard-backend/internals/tickboard/products/users/user"
)

// NewApp creates and configures the main application.
func NewApp() application.Application {
	return application.NewBuilder().
		WithOptions(
			config.Module,
			logger.Module,
			pgsql.Module,
			rds.Module,
			httpx.Module,
		).WithOptions(
		user.Module,
	).Build()
}

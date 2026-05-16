package tickboard

import (
	"context"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/config"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/httpx"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/logger"
	rds "github.com/halimdotnet/tickboard-backend/internals/pkg/redis"
	"go.uber.org/fx"
)

type App struct {
	app *fx.App
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
	a.app = fx.New(
		fx.Options(
			config.Module,
			logger.Module,
			rds.Module,
			httpx.Module,
		),
	)
	a.app.Run()
}

func (a *App) Start(ctx context.Context) error {
	return a.app.Start(ctx)
}

func (a *App) Stop(ctx context.Context) error {
	return a.app.Stop(ctx)
}

func (a *App) Wait() <-chan fx.ShutdownSignal {
	return a.app.Wait()
}

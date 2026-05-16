package application

import (
	"context"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/config"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

type builder struct {
	options    []fx.Option
	providers  []interface{}
	invokes    []interface{}
	decorators []interface{}
}

// NewBuilder creates a new application builder.
func NewBuilder() Builder {
	return &builder{
		options:    make([]fx.Option, 0),
		providers:  make([]interface{}, 0),
		invokes:    make([]interface{}, 0),
		decorators: make([]interface{}, 0),
	}
}

func (b *builder) WithOptions(opts ...fx.Option) Builder {
	b.options = append(b.options, opts...)
	return b
}

func (b *builder) WithProviders(constructors ...interface{}) Builder {
	b.providers = append(b.providers, constructors...)
	return b
}

func (b *builder) WithInvokes(funcs ...interface{}) Builder {
	b.invokes = append(b.invokes, funcs...)
	return b
}

func (b *builder) WithDecorators(decorators ...interface{}) Builder {
	b.decorators = append(b.decorators, decorators...)
	return b
}

func (b *builder) Build() Application {
	opts := append([]fx.Option{}, b.options...)

	if len(b.providers) > 0 {
		opts = append(opts, fx.Provide(b.providers...))
	}
	if len(b.invokes) > 0 {
		opts = append(opts, fx.Invoke(b.invokes...))
	}
	if len(b.decorators) > 0 {
		opts = append(opts, fx.Decorate(b.decorators...))
	}

	opts = append(opts, fx.WithLogger(func(log logger.Logger, cfg *config.Config) fxevent.Logger {
		if !cfg.Debug {
			return fxevent.NopLogger
		}
		return &fxLogger{log: log}
	}))

	app := fx.New(opts...)
	return &application{app: app}
}

type application struct {
	app *fx.App
}

func (a *application) Run() {
	a.app.Run()
}

func (a *application) Start(ctx context.Context) error {
	return a.app.Start(ctx)
}

func (a *application) Stop(ctx context.Context) error {
	return a.app.Stop(ctx)
}

func (a *application) Wait() <-chan fx.ShutdownSignal {
	return a.app.Wait()
}

package application

import (
	"context"

	"go.uber.org/fx"
)

// Builder defines the interface for creating a new Application.
type Builder interface {
	WithOptions(opts ...fx.Option) Builder
	WithProviders(constructors ...interface{}) Builder
	WithInvokes(funcs ...interface{}) Builder
	WithDecorators(decorators ...interface{}) Builder
	Build() Application
}

// Application defines the interface for the running application.
type Application interface {
	Run()
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Wait() <-chan fx.ShutdownSignal
}

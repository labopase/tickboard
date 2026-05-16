package healthz

import (
	"context"

	"github.com/labstack/echo/v5"
)

const (
	StatusOk    = "ok"
	StatusError = "error"
)

type CheckResult struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type CheckFunc func(ctx context.Context) error

// Checker defines the contract for managing and executing health checks.
type Checker interface {
	// Register adds a new health check with the given name.
	Register(name string, fn CheckFunc)
	// Check runs all registered health checks concurrently and returns the results.
	Check(ctx context.Context) map[string]CheckResult
}

type Handler interface {
	LivenessHandler(c *echo.Context) error
	ReadinessHandler(c *echo.Context) error
	StartupHandler(c *echo.Context) error
	SetReady(ready bool)
}

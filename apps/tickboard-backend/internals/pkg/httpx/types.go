package httpx

import (
	"context"

	"github.com/labstack/echo/v5"
)

type Engine interface {
	Start(ctx context.Context) error
	Instance() *echo.Echo
}

type Response struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

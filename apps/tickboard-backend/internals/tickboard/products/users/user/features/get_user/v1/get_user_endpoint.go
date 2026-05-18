package v1

import (
	"net/http"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/httpx"
	"github.com/halimdotnet/tickboard-backend/internals/pkg/logger"
	"github.com/labstack/echo/v5"
)

type getUserEndpoint struct {
	log  logger.Logger
	echo httpx.Engine
}

type GetUserEndpoint interface {
	RegisterRoute()
}

func NewGetUserEndpoint(log logger.Logger, echo httpx.Engine) GetUserEndpoint {
	return &getUserEndpoint{log: log, echo: echo}
}

func (e *getUserEndpoint) RegisterRoute() {
	e.echo.Instance().GET("/data", e.getUserByID)
}

func (e *getUserEndpoint) getUserByID(c *echo.Context) error {
	return c.JSON(http.StatusOK, httpx.Response{
		Error:   false,
		Message: "success",
		Data:    nil,
	})
}

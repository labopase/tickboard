package httpx

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/halimdotnet/tickboard-backend/internals/pkg/logger"
	"github.com/labstack/echo/v5"
)

type engine struct {
	echo   *echo.Echo
	log    logger.Logger
	config *Config
}

func NewEngine(cfg *Config, log logger.Logger) (Engine, error) {
	if cfg == nil {
		return nil, fmt.Errorf("httpx: config is required")
	}

	cfg.applyDefaults()

	e := echo.New()

	return &engine{
		echo:   e,
		log:    log,
		config: cfg,
	}, nil
}

func (e *engine) Start(ctx context.Context) error {
	ln, err := net.Listen("tcp", e.config.Addr())
	if err != nil {
		return fmt.Errorf("httpx: failed to listen on %s: %w", e.config.Addr(), err)
	}

	e.log.Infof("starting http server on %s", e.config.Addr())

	sc := echo.StartConfig{
		Address:         e.config.Addr(),
		Listener:        ln,
		HideBanner:      true,
		HidePort:        true,
		GracefulTimeout: e.config.ShutdownTimeout,
		BeforeServeFunc: func(s *http.Server) error {
			s.ReadTimeout = e.config.ReadTimeout
			s.IdleTimeout = e.config.IdleTimeout
			s.MaxHeaderBytes = e.config.MaxHeaderBytes
			return nil
		},
		OnShutdownError: func(err error) {
			e.log.Errorf("http server error on shutdown %s", err.Error())
		},
	}

	return sc.Start(ctx, e.echo)
}

func (e *engine) Instance() *echo.Echo {
	return e.echo
}

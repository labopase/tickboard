package healthz

import (
	"net/http"
	"sync/atomic"

	"github.com/labstack/echo/v5"
)

type handler struct {
	checker Checker
	ready   *atomic.Bool
}

func NewHandler(checker Checker) Handler {

	return &handler{
		checker: checker,
		ready:   &atomic.Bool{},
	}
}

func (h *handler) LivenessHandler(c *echo.Context) error {
	return c.JSON(http.StatusOK, CheckResult{
		Status: StatusOk,
	})
}

func (h *handler) ReadinessHandler(c *echo.Context) error {
	if !h.ready.Load() {
		return c.JSON(http.StatusServiceUnavailable, CheckResult{
			Status: StatusError,
			Error:  "service not ready",
		})
	}

	results := h.checker.Check(c.Request().Context())

	for _, result := range results {
		if result.Status == StatusError {
			return c.JSON(http.StatusServiceUnavailable, CheckResult{
				Status: StatusError,
				Error:  "service not ready",
			})
		}
	}

	return c.JSON(http.StatusOK, CheckResult{
		Status: StatusOk,
	})
}

func (h *handler) StartupHandler(c *echo.Context) error {
	if h.ready.Load() {
		return c.JSON(http.StatusOK, CheckResult{
			Status: StatusOk,
		})
	} else {
		return c.JSON(http.StatusServiceUnavailable, CheckResult{
			Status: StatusError,
			Error:  "service unavailable",
		})
	}
}

func (h *handler) SetReady(ready bool) {
	h.ready.Store(ready)
}

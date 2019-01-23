package healthcheck

import (
	"net/http"
	"sync/atomic"
	"time"

	"github.com/arifsetiawan/nextpkg/env"
	"github.com/labstack/echo"
)

// Handler ...
type Handler struct {
	IsReady *atomic.Value
}

// NewHealthCheckHandler ...
func NewHealthCheckHandler() *Handler {
	isReady := &atomic.Value{}
	isReady.Store(false)

	return &Handler{
		IsReady: isReady,
	}
}

// SetRoutes ...
func (h *Handler) SetRoutes(r *echo.Echo) {
	r.GET("/live", h.live)
	r.GET("/ready", h.ready, h.readyMiddleware(h.IsReady))
	r.GET("/sleep", h.sleep)
}

// Ready ...
func (h *Handler) Ready(ready bool) {
	h.IsReady.Store(ready)
}

func (h *Handler) live(c echo.Context) error {
	return c.String(http.StatusOK, "app: "+env.Getenv("APP_NAME", "service")+", v: "+env.Getenv("VERSION", "0.1.0"))
}

func (h *Handler) ready(c echo.Context) error {
	return c.String(http.StatusOK, "app: "+env.Getenv("APP_NAME", "service")+", v: "+env.Getenv("VERSION", "0.1.0"))
}

func (h *Handler) readyMiddleware(isReady *atomic.Value) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if h.IsReady == nil || !h.IsReady.Load().(bool) {
				return c.NoContent(http.StatusServiceUnavailable)
			}
			return next(c)
		}
	}
}

func (h *Handler) sleep(c echo.Context) error {
	time.Sleep(5 * time.Second)
	return c.String(http.StatusOK, "app: "+env.Getenv("APP_NAME", "service")+", v: "+env.Getenv("VERSION", "0.1.0"))
}

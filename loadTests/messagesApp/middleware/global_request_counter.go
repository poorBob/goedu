package middleware

import (
	"sync/atomic"

	"github.com/labstack/echo/v4"
)

type GlobalRequestCounter struct {
	totalRequests uint64
}

func NewGlobalRequestCounter() *GlobalRequestCounter {
	return &GlobalRequestCounter{}
}

func (g *GlobalRequestCounter) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			atomic.AddUint64(&g.totalRequests, 1)
			return next(c)
		}
	}
}

func (g *GlobalRequestCounter) TotalRequests() uint64 {
	return atomic.LoadUint64(&g.totalRequests)
}

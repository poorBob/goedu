package middleware

import (
	"sync/atomic"

	"github.com/labstack/echo/v4"
)

type SpecificRequestStats struct {
	PostCount int64
	GetCount  int64
}

func NewSpecificRequestStatsMiddleware(stats *SpecificRequestStats) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Method == "POST" && c.Path() == "/api/message" {
				atomic.AddInt64(&stats.PostCount, 1)
			} else if c.Request().Method == "GET" && c.Path() == "/api/message" {
				atomic.AddInt64(&stats.GetCount, 1)
			}
			return next(c)
		}
	}
}

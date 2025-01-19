package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

func NewRateLimiterMiddleware(rps rate.Limit, burst int) echo.MiddlewareFunc {
	limiter := rate.NewLimiter(rps, burst)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if limiter.Allow() {
				return next(c)
			}

			return c.String(http.StatusTooManyRequests, "Too Many Requests")
		}
	}
}

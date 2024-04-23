package internal

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

// Create a new rate limiter instance
var limiter = rate.NewLimiter(1, 5) // 1 request per second with a burst of 5

// RateLimiterMiddleware checks if the request exceeds the rate limit
func RateLimiterMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !limiter.Allow() {
			return echo.NewHTTPError(http.StatusTooManyRequests, "Too many requests")
		}
		return next(c)
	}
}

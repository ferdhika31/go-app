package middleware

import (
	"github.com/didip/tollbooth/v7"
	"time"

	"github.com/didip/tollbooth/v7/limiter"
	"github.com/labstack/echo/v4"
)

func Limiter(next echo.HandlerFunc) echo.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Hour,
	})

	lmt.SetIPLookups([]string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"})
	lmt.SetMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"})

	return func(c echo.Context) error {
		httpError := tollbooth.LimitByRequest(lmt, c.Response(), c.Request())

		if httpError != nil {
			return &echo.HTTPError{
				Code:    httpError.StatusCode,
				Message: "Too many requests, please try again later.",
			}
		}

		return next(c)
	}
}

package inertiaMiddleware

import "github.com/labstack/echo/v4"

func InertiaMiddleware(e *echo.Echo) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	}
}

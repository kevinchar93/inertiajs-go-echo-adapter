package main

import (
	"inertia-echo/inertiaMiddleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func configureMiddleware(e *echo.Echo) {
	e.Use(inertiaMiddleware.InertiaMiddleware(e))
}

func configureRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}

func main() {
	e := echo.New()

	configureMiddleware(e)
	configureRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}

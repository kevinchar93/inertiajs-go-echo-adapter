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
	e.GET("/", func(ctx echo.Context) error {

		data := map[string]interface{}{}

		return ctx.Render(http.StatusOK, "Index", data)
	})
}

func main() {
	e := echo.New()

	configureMiddleware(e)
	configureRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}

package main

import (
	"inertia-echo/middleware"
	"inertia-echo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	middleware.ConfigureMiddleware(e)
	routes.ConfigureRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}

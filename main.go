package main

import (
	"inertia-echo/inertiaMiddleware"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func configureMiddleware(e *echo.Echo) {

	assetVersion := "1"
	e.Use(inertiaMiddleware.InertiaMiddleware(e, assetVersion))

	e.Static("/", "views/public")

	isDevelopment := os.Getenv("BUILD_ENV") == "development"

	if isDevelopment {
		viteDevServer, err := url.Parse("http://localhost:5173")
		if err != nil {
			log.Fatal("Could not parse Vite dev server url", err)
		}

		devAssets := e.Group("/src/assets")
		target := []*middleware.ProxyTarget{
			{Name: "viteProxyTarget",
				URL: viteDevServer}}

		loadbalancer := middleware.NewRoundRobinBalancer(target)

		devAssets.Use(middleware.Proxy(loadbalancer))
		return
	}
}

func configureRoutes(e *echo.Echo) {
	e.GET("/", func(ctx echo.Context) error {

		props := map[string]interface{}{
			"exampleProp": "Let's a go!",
		}

		return ctx.Render(http.StatusOK, "Index", props)
	})

	e.GET("/example-page", func(ctx echo.Context) error {

		props := map[string]interface{}{
			"phrase": "Don't panic!",
		}

		return ctx.Render(http.StatusOK, "ExamplePage", props)
	})
}

func main() {
	e := echo.New()

	configureMiddleware(e)
	configureRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}

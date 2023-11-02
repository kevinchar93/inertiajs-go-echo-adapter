package routes

import (
	"inertia-echo/page"
	"inertia-echo/pokemondb"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ConfigureRoutes(e *echo.Echo, db map[int]pokemondb.PokemonInfo) {
	e.GET("/", func(ctx echo.Context) error {

		header := page.NewPageHeader("Index title", "meta name", "meta content")

		props := map[string]interface{}{
			"exampleProp": "Let's a go!",
			"header":      header,
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

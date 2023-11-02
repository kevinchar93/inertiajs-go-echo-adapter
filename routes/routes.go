package routes

import (
	"cmp"
	"inertia-echo/page"
	"inertia-echo/pokemondb"
	"net/http"

	"slices"

	"github.com/labstack/echo/v4"
)

func ConfigureRoutes(e *echo.Echo, db map[int]pokemondb.PokemonInfo) {
	e.GET("/", func(ctx echo.Context) error {

		header := page.NewPageHeader("Index title", "meta name", "meta content")

		type PokemonListItem struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		var allPokemon []PokemonListItem

		for _, pokemon := range db {
			allPokemon = append(allPokemon, PokemonListItem{ID: pokemon.ID, Name: pokemon.Name.English})
		}

		slices.SortFunc(allPokemon, func(a, b PokemonListItem) int {
			return cmp.Compare(a.ID, b.ID)
		})

		props := map[string]interface{}{
			"header":     header,
			"allPokemon": allPokemon,
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

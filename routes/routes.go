package routes

import (
	"cmp"
	"inertia-echo/page"
	"inertia-echo/pokemondb"
	"net/http"
	"strconv"

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

	e.GET("/pokemon/:id", func(ctx echo.Context) error {

		id, _ := strconv.Atoi(ctx.Param("id"))
		pokemon := db[id]
		name := pokemon.Name.English

		clampNextOrPrevPokemon := func(id int) int {
			if id <= 0 || id >= len(db)+1 {
				return -1
			}
			return id
		}

		header := page.NewPageHeader(name, "Pages for "+name, "Pages with stats & info about "+name)

		props := map[string]interface{}{
			"header":      header,
			"number":      id,
			"name":        name,
			"nextPokemon": clampNextOrPrevPokemon(id + 1),
			"prevPokemon": clampNextOrPrevPokemon(id - 1),
			"type":        pokemon.Type,
			"stats":       pokemon.Base,
		}

		return ctx.Render(http.StatusOK, "Pokemon", props)
	})
}

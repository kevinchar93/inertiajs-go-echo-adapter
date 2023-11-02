package main

import (
	"inertia-echo/middleware"
	"inertia-echo/pokemondb"
	"inertia-echo/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	middleware.ConfigureMiddleware(e)

	db := pokemondb.CreatePokemonDB()
	routes.ConfigureRoutes(e, db)

	e.Logger.Fatal(e.Start(":3000"))
}

package pokemondb

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type PokemonInfo struct {
	ID int `json:"id"`

	Name struct {
		English  string `json:"english"`
		Japanese string `json:"japanese"`
		Chinese  string `json:"chinese"`
		French   string `json:"french"`
	} `json:"name"`

	Type []string `json:"type"`

	Base struct {
		Hp        int `json:"hp"`
		Attack    int `json:"attack"`
		Defense   int `json:"defense"`
		SpAttack  int `json:"spAttack"`
		SpDefense int `json:"spDefense"`
		Speed     int `json:"speed"`
	} `json:"base"`
}

func CreatePokemonDB() map[int]PokemonInfo {
	jsonFile, err := os.Open("./pokemondb/pokedex.json")
	if err != nil {
		log.Fatal("Failed to open pokedex.json")
	}
	defer jsonFile.Close()

	jsonBytes, _ := io.ReadAll(jsonFile)
	var pokedexData []PokemonInfo
	json.Unmarshal([]byte(jsonBytes), &pokedexData)

	pokemonDb := make(map[int]PokemonInfo, len(pokedexData))

	for _, pokemonInfo := range pokedexData {
		pokemonDb[pokemonInfo.ID] = pokemonInfo
	}

	return pokemonDb
}

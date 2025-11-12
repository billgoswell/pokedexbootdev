package main

import (
	"github.com/billgoswell/pokedexbootdev/internal/pokeapi"
	"time"
)

type config struct {
	pokedex             map[string]pokeapi.PokemonResp
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokedex:       make(map[string]pokeapi.PokemonResp),
		pokeapiClient: pokeapi.NewClient(time.Minute),
	}
	startRepl(&cfg)
}

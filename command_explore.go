package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No Location Area provided")
	}
	resp, err := cfg.pokeapiClient.LocationArea(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("%s information:\n", args[0])
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}

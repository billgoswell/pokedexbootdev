package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		fmt.Println("You have no pokemon in your pokedex")
		return nil
	}
	for pokemon, _ := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}

package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No Pokemon provided")
	}
	pokemon, err := cfg.pokeapiClient.Pokemon(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	chance := rand.Intn(pokemon.BaseExperience)
	if chance < 100 {
		fmt.Printf("You caught %s\n", args[0])
		cfg.pokedex[args[0]] = pokemon
		return nil
	}
	fmt.Printf("%s escaped capture", args[0])
	return nil
}

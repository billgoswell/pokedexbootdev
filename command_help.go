package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Here are your available commands: ")
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("- %s: %s \n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "See help",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "list locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "list last page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "list pokemon in a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Returns info about a Pokemon if you have caught it",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "catch",
			description: "Attempts to catch the pokemon",
			callback:    commandPokedex,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		fmt.Printf("Your command was: %s\n", words[0])
		command, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unknown Command")
			continue
		}
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

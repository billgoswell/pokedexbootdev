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
	callback    func(*config) error
}
type config struct {
	next     string
	previous string
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
	}
}

func main() {
	config := config{next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20", previous: ""}
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
		command.callback(&config)
	}
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: \n")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	if len(text) == 0 {
		return []string{}
	}
	text = strings.ToLower(text)
	words := strings.Split(text, " ")
	return words
}

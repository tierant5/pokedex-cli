package main

import "github.com/tierant5/pokedex-cli/internal/pokeapi"

type cliCommand struct {
	callback    func(config *config, params []string) error
	name        string
	description string
}

type config struct {
	pokeapiClient pokeapi.Client
	Pokedex       *Pokedex
	Next          string
	Previous      string
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
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area-name>",
			description: "Explore the <area-name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Try to catch <pokemon-name>",
			callback:    commandCatch,
		},
	}
}

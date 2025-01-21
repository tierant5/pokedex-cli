package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandHelp() error {
	commandMap := map[string]cliCommand{
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
	}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range commandMap {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	callback    func() error
	name        string
	description string
}

func main() {
	commandMap := map[string]cliCommand{
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
	}
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "
	for {
		fmt.Print(prompt)
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		cleanedInput := cleanInput(text)
		command, ok := commandMap[cleanedInput[0]]
		if !ok {
			fmt.Printf("unknown command: %s\n", cleanedInput[0])
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Printf("something went wrong...\n%v", fmt.Errorf("%w", err))
		}
	}
}

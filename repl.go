package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL(cfg *config) {
	commandMap := getCommands()
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
		err := command.callback(cfg, cleanedInput[1:])
		if err != nil {
			fmt.Printf("something went wrong...\n%v", fmt.Errorf("%w", err))
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

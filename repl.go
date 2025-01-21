package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	commandMap := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "
	config := config{}
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
		err := command.callback(&config)
		if err != nil {
			fmt.Printf("something went wrong...\n%v", fmt.Errorf("%w", err))
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

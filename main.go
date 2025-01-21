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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "
	for {
		fmt.Print(prompt)
		scanner.Scan()
		text := scanner.Text()
		cleaned_text := cleanInput(text)
		fmt.Printf("Your command was: %s\n", cleaned_text[0])
	}
}

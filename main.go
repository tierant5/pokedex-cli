package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Split(strings.TrimSpace(text), " ")
}

func main() {
	fmt.Println("Hello, World!")
}

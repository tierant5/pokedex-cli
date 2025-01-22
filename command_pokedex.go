package main

import (
	"fmt"
)

func commandPokedex(config *config, params []string) error {
	fmt.Println("Your Pokedex:")
	names := config.Pokedex.GetAllPokemon()
	for _, name := range names {
		fmt.Printf("  - %v\n", name)
	}
	return nil
}

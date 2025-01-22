package main

import (
	"fmt"
)

func commandInspect(config *config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("no <pokemon-name> was provided\n")
	}
	pokemonName := params[0]
	pokemonInfo, ok := config.Pokedex.GetPokemon(pokemonName)
	if !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		pokemonInfo.PrintInfo()
	}
	return nil
}

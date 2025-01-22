package main

import (
	"fmt"
	"math/rand"

	"github.com/tierant5/pokedex-cli/internal/pokeapi"
)

func commandCatch(config *config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("no <pokemon-name> was provided\n")
	}
	pokemonName := params[0]
	url := pokeapi.BaseUrl + "pokemon/" + pokemonName
	pokemonInfo, err := pokeapi.GetPokeAPIType[pokeapi.PokemonInfo](&config.pokeapiClient, url)
	if err != nil {
		return err
	}
	chance := int((float64(rand.Intn(pokemonInfo.BaseExperience)) / float64(pokemonInfo.BaseExperience)) * 100)
	roll := rand.Intn(100)
	caught := roll <= chance
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonName)
	if caught {
		fmt.Printf("%v was caught!\n", pokemonName)
		config.Pokedex.AddPokemon(pokemonInfo)
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}
	return nil
}

package main

import (
	"fmt"

	"github.com/tierant5/pokedex-cli/internal/pokeapi"
)

func commandExplore(config *config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("no <area-name> was provided\n")
	}
	areaName := params[0]
	url := pokeapi.BaseUrl + "location-area/" + areaName
	areaInfo, err := pokeapi.GetPokeAPIType[pokeapi.LocationAreaInfo](&config.pokeapiClient, url)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", areaName)
	areaInfo.PrintPokemon()
	return nil
}

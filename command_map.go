package main

import (
	"fmt"

	"github.com/tierant5/pokedex-cli/internal/pokeapi"
)

func commandMap(config *config, params []string) error {
	page, err := pokeapi.GetPokeAPIType[pokeapi.LocationAreaPage](&config.pokeapiClient, config.Next)
	if err != nil {
		return err
	}
	if page.Next != nil {
		config.Next = *page.Next
	} else {
		config.Next = ""
	}
	if page.Previous != nil {
		config.Previous = *page.Previous
	} else {
		config.Previous = ""
	}
	page.PrintAreas()
	return nil
}

func commandMapb(config *config, params []string) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	page, err := pokeapi.GetPokeAPIType[pokeapi.LocationAreaPage](&config.pokeapiClient, config.Previous)
	if err != nil {
		return err
	}
	if page.Next != nil {
		config.Next = *page.Next
	} else {
		config.Next = ""
	}
	if page.Previous != nil {
		config.Previous = *page.Previous
	} else {
		config.Previous = ""
	}
	page.PrintAreas()
	return nil
}

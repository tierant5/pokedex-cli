package main

import (
	"fmt"
	"internal/pokeapi"
)

func commandMap(config *config) error {
	page, err := pokeapi.GetLocationAreaPage(config.Next)
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

func commandMapb(config *config) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	page, err := pokeapi.GetLocationAreaPage(config.Previous)
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

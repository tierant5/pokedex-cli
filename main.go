package main

import (
	"time"

	"github.com/tierant5/pokedex-cli/internal/pokeapi"
)

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
	}
	startREPL(&cfg)
}

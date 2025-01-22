package pokeapi

import "fmt"

type LocationAreaInfo struct {
	Location          Location            `json:"location"`
	Name              string              `json:"name"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
	ID                int                 `json:"id"`
	GameIndex         int                 `json:"game_index"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

func (l LocationAreaInfo) GetApiResource() string {
	return "location-area/"
}

func (l LocationAreaInfo) PrintPokemon() {
	fmt.Println("Found Pokemon:")
	for _, encounter := range l.PokemonEncounters {
		fmt.Printf("  - %v\n", encounter.Pokemon.Name)
	}
}

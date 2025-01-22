package pokeapi

import "fmt"

type PokemonInfo struct {
	Name           string  `json:"name"`
	Stats          []Stats `json:"stats"`
	Types          []Types `json:"types"`
	ID             int     `json:"id"`
	BaseExperience int     `json:"base_experience"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
}
type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Stats struct {
	Stat     Stat `json:"stat"`
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
}
type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Types struct {
	Type Type `json:"type"`
	Slot int  `json:"slot"`
}

func (p PokemonInfo) GetApiResource() string {
	return "pokemon/"
}

func (p PokemonInfo) PrintInfo() {
	fmt.Printf("Name: %v\n", p.Name)
	fmt.Printf("Height: %v\n", p.Height)
	fmt.Printf("Weight: %v\n", p.Weight)
	fmt.Println("Stats:")
	for _, stat := range p.Stats {
		fmt.Printf("  - %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, ptype := range p.Types {
		fmt.Printf("  - %v\n", ptype.Type.Name)
	}
}

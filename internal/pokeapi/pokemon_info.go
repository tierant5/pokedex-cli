package pokeapi

type PokemonInfo struct {
	Name           string `json:"name"`
	ID             int    `json:"id"`
	BaseExperience int    `json:"base_experience"`
}

func (p PokemonInfo) GetApiResource() string {
	return "pokemon/"
}

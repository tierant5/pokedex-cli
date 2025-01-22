package main

import "github.com/tierant5/pokedex-cli/internal/pokeapi"

type Pokedex struct {
	pokemon map[string]pokeapi.PokemonInfo
}

func NewPokedex() *Pokedex {
	return &Pokedex{pokemon: map[string]pokeapi.PokemonInfo{}}
}

func (p *Pokedex) AddPokemon(pokemon pokeapi.PokemonInfo) error {
	p.pokemon[pokemon.Name] = pokemon
	return nil
}

func (p *Pokedex) GetPokemon(name string) (pokeapi.PokemonInfo, bool) {
	pokemonInfo, ok := p.pokemon[name]
	return pokemonInfo, ok
}

func (p *Pokedex) GetAllPokemon() []string {
	if len(p.pokemon) == 0 {
		return []string{}
	}
	var pokemon []string
	for name := range p.pokemon {
		pokemon = append(pokemon, name)
	}
	return pokemon
}

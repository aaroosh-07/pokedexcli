package pokedex

import "github.com/aaroosh-07/pokedexcli/internal/pokeapi"

type PokedexStruct struct {
	pokedexInfo map[string]pokeapi.PokemonData
}

func NewPokedex() (*PokedexStruct) {
	return &PokedexStruct{pokedexInfo: make(map[string]pokeapi.PokemonData)}
}

func (p *PokedexStruct) Add(pokename string, data pokeapi.PokemonData) bool {
	_, isPresent := p.pokedexInfo[pokename]

	if isPresent {
		return false
	}

	p.pokedexInfo[pokename] = data

	return true
}

func (p *PokedexStruct) Get(pokename string) (pokeapi.PokemonData, bool) {
	_, isPresent := p.pokedexInfo[pokename]

	if !isPresent {
		return pokeapi.PokemonData{}, false
	}

	return p.pokedexInfo[pokename], true
}
package pokedex

import (
	"fmt"

	"github.com/aaroosh-07/pokedexcli/internal/pokeapi"
)

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

func (p *PokedexStruct) DisplayPokeInfo(pokename string) error {
	data, isPresent := p.pokedexInfo[pokename]

	if !isPresent {
		return fmt.Errorf("pokemon not present in pokedex")
	}

	var pokeInfoStr string

	pokeInfoStr = fmt.Sprintf("Name: %s\nHeight: %v\nWeight: %v\n", data.Name, data.Height, data.Weight)

	var statStr string
	statStr = fmt.Sprintf("Stats:\n")

	for _, stat := range data.Stats {
		statStr += fmt.Sprintf("\t-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	var typeStr string
	typeStr = fmt.Sprintf("Types:\n")

	for _, t := range data.Types {
		typeStr += fmt.Sprintf("\t- %s\n", t.Type.Name)
	}

	pokeInfoStr += (statStr + typeStr)
	fmt.Print(pokeInfoStr)
	return nil
}

func (p *PokedexStruct) GetNumPokemons() int {
	return len(p.pokedexInfo)
}

func (p *PokedexStruct) GetPokemonName() []string {
	var pokenames = []string{}

	for key := range p.pokedexInfo {
		pokenames = append(pokenames, key)
	}

	return pokenames
}
package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/aaroosh-07/pokedexcli/internal/pokecache"
)

func FetchPokemonData(name string, c *pokecache.Cache) (*PokemonData, error) {
	fullUrl := baseUrl + fmt.Sprintf("/pokemon/%s", name)

	var resData []byte
	var err error
	resData, isPresent := c.Get(fullUrl)

	if !isPresent {
		resData, err = fetchApiData(fullUrl)

		if err != nil {
			return &PokemonData{}, err
		}

		//add data to cache if not found
		c.Add(fullUrl, resData)
	}

	//unmarshal the []byte to go struct

	var pokeData PokemonData
	err = json.Unmarshal(resData, &pokeData)

	if err != nil {
		return &PokemonData{}, err
	}

	return &pokeData, nil
}
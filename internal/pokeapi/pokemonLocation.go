package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/aaroosh-07/pokedexcli/internal/pokecache"
)


type locationData struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetPokemonAtLocation(location string, c *pokecache.Cache) ([]string, error) {
	fullUrl := fmt.Sprintf("%s/location-area/%s", baseUrl, location)

	var resData []byte
	var err error

	resData, isPresent := c.Get(fullUrl)

	if !isPresent {
		resData, err = fetchApiData(fullUrl)
		//update cache
		if err != nil {
			return []string{}, err
		}

		c.Add(fullUrl, resData)
	}

	//read data into go struct
	var jsonData locationData

	err = json.Unmarshal(resData, &jsonData)

	if err != nil {
		return []string{}, fmt.Errorf("error unpacking unicode data to go struct")
	}

	pokemonlist := fetchPokemonNames(&jsonData)

	return pokemonlist, nil
}

func fetchPokemonNames(info *locationData) []string {
	var result []string

	for _, pokeData := range info.PokemonEncounters {
		result = append(result, pokeData.Pokemon.Name)
	}

	return result
}

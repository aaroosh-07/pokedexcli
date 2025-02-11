package pokeapi

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aaroosh-07/pokedexcli/internal/pokecache"
)

type PokeApiLocation struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []struct{
		Name string `json:"name"`
		Url string `json:"url"`
	} `json:"results"`
}

func GetPokeapiLocation(limit, offset int, cache *pokecache.Cache) (PokeApiLocation , error) {
	var pokeapi_url string = baseUrl + "/location-area"

	fullUrl := fmt.Sprintf("%s?limit=%s&offset=%s",pokeapi_url, strconv.Itoa(limit), strconv.Itoa(offset))

	var resData []byte
	var err error

	//check if data present in cache
	resData, isPresent := cache.Get(fullUrl)

	//if not present in cache then make network request
	if !isPresent {
		resData, err = fetchApiData(fullUrl)
	
		if err != nil {
			return PokeApiLocation{}, err
		}

		//add data in cache
		cache.Add(fullUrl, resData)
	}

	var resJsonData PokeApiLocation

	err = json.Unmarshal(resData, &resJsonData)

	if err != nil {
		return PokeApiLocation{}, fmt.Errorf("error converting json to go struct")
	}

	return resJsonData, nil
}

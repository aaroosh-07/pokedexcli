package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://pokeapi.co/api/v2"

func fetchApiData(fullUrl string) ([]byte, error) {
	res, err := http.Get(fullUrl)

	if err != nil {
		return []byte{}, fmt.Errorf("network error: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("bad response status code: %d", res.StatusCode)
	}

	resData, err := io.ReadAll(res.Body)

	if err != nil {
		return []byte{}, fmt.Errorf("error reading data")
	}

	return resData, nil
}
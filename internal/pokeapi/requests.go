package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tierant5/pokedex-cli/internal/pokecache"
)

func GetLocationAreaPage(url string, cache *pokecache.Cache) (LocationAreaPage, error) {
	if url == "" {
		baseUrl := "https://pokeapi.co/api/v2/"
		fullUrl := baseUrl + "location-area/"
		url = fullUrl
	}

	// Check if the url is in the cache, return the cached data early
	cached_data, ok := cache.Get(url)
	var locationAreaPage LocationAreaPage
	if ok {
		err := json.Unmarshal(cached_data, &locationAreaPage)
		if err == nil {
			return locationAreaPage, nil
		} else {
			fmt.Printf("error recreating cached data: %v", err)
		}
	}

	// TODO: Split this into it's own function
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaPage{}, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return LocationAreaPage{}, err
	}
	if resp.StatusCode > 299 {
		return LocationAreaPage{}, fmt.Errorf("Bad HTTP Request: %v", resp.Status)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaPage{}, err
	}

	err = json.Unmarshal(data, &locationAreaPage)
	if err != nil {
		return LocationAreaPage{}, err
	}

	cache.Add(url, data)

	return locationAreaPage, nil
}

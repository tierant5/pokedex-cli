package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetLocationAreaPage(url string) (LocationAreaPage, error) {
	if url == "" {
		baseUrl := "https://pokeapi.co/api/v2/"
		fullUrl := baseUrl + "location-area/"
		url = fullUrl
	}
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

	var locationAreaPage LocationAreaPage
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locationAreaPage)
	if err != nil {
		return LocationAreaPage{}, err
	}

	return locationAreaPage, nil
}

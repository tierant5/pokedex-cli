package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaPage(url string) (LocationAreaPage, error) {
	if url == "" {
		fullUrl := baseUrl + "location-area/"
		url = fullUrl
	}

	// Check if the url is in the cache, return the cached data early
	cached_data, ok := c.cache.Get(url)
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
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaPage{}, err
	}
	if resp.StatusCode > 299 {
		return LocationAreaPage{}, fmt.Errorf("Bad HTTP Request: %v\n", resp.Status)
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

	c.cache.Add(url, data)

	return locationAreaPage, nil
}

package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetPokeAPIType[T PokeAPIType](c *Client, url string) (T, error) {
	var zeroT T
	var pokeApiType T
	if url == "" {
		fullUrl := BaseUrl + pokeApiType.getApiResource()
		url = fullUrl
	}

	// Check if the url is in the cache, return the cached data early
	cached_data, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(cached_data, &pokeApiType)
		if err == nil {
			return pokeApiType, nil
		} else {
			fmt.Printf("error recreating cached data: %v", err)
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return zeroT, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return zeroT, err
	}
	if resp.StatusCode > 299 {
		return zeroT, fmt.Errorf("Bad HTTP Request: %v\n", resp.Status)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return zeroT, err
	}

	err = json.Unmarshal(data, &pokeApiType)
	if err != nil {
		return zeroT, err
	}

	c.cache.Add(url, data)

	return pokeApiType, nil
}

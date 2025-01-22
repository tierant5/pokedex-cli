package pokeapi

import "fmt"

type LocationAreaPage struct {
	Previous *string               `json:"previous"`
	Next     *string               `json:"next"`
	Results  []LocationAreaResults `json:"results"`
	Count    int                   `json:"count"`
}
type LocationAreaResults struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (l LocationAreaPage) PrintAreas() {
	for _, result := range l.Results {
		fmt.Println(result.Name)
	}
}

func (l LocationAreaPage) getApiResource() string {
	return "location-area/"
}

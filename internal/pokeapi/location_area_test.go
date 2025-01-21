package pokeapi

import (
	"encoding/json"
	"testing"
)

func TestLocationArea(t *testing.T) {
	testNextUrl := "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20"
	cases := []struct {
		input    string
		expected LocationAreaPage
	}{
		{
			input: `{
  "count": 1054,
  "next": "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
  "previous": null,
  "results": [
    {
      "name": "canalave-city-area",
      "url": "https://pokeapi.co/api/v2/location-area/1/"
    },
    {
      "name": "eterna-city-area",
      "url": "https://pokeapi.co/api/v2/location-area/2/"
    }
  ]
}`,
			expected: LocationAreaPage{
				Count:    1054,
				Next:     &testNextUrl,
				Previous: nil,
				Results: []LocationAreaResults{
					{
						Name: "canalave-city-area",
						URL:  "https://pokeapi.co/api/v2/location-area/1/",
					},
					{
						Name: "eterna-city-area",
						URL:  "https://pokeapi.co/api/v2/location-area/2/",
					},
				},
			},
		},
	}
	for _, c := range cases {
		var locationAreaPage LocationAreaPage
		err := json.Unmarshal([]byte(c.input), &locationAreaPage)
		if err != nil {
			t.Errorf("error trying to unmarshal json: %v", err)
		}
		if locationAreaPage.Count != c.expected.Count {
			t.Errorf("%v != %v", locationAreaPage.Count, c.expected.Count)
		}
		if locationAreaPage.Next != nil && c.expected.Next != nil {
			if *locationAreaPage.Next != *c.expected.Next {
				t.Errorf("%v != %v", locationAreaPage.Next, c.expected.Next)
			}
		} else {
			if locationAreaPage.Next != nil || c.expected.Next != nil {
				t.Error("both locationAreaPage.Next and c.expected.Next are not nil")
			}
		}
		if locationAreaPage.Previous != nil && c.expected.Previous != nil {
			if *locationAreaPage.Previous != *c.expected.Previous {
				t.Errorf("%v != %v", locationAreaPage.Previous, c.expected.Previous)
			}
		} else {
			if locationAreaPage.Previous != nil || c.expected.Previous != nil {
				t.Error("both locationAreaPage.Previous and c.expected.Previous are not nil")
			}
		}
		for i, result := range c.expected.Results {
			if locationAreaPage.Results[i].Name != result.Name {
				t.Errorf("%v != %v", locationAreaPage.Results[i].Name, result.Name)
			}
			if locationAreaPage.Results[i].URL != result.URL {
				t.Errorf("%v != %v", locationAreaPage.Results[i].URL, result.URL)
			}
		}
	}
}

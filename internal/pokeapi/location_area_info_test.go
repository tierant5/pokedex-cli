package pokeapi

import (
	"encoding/json"
	"testing"
)

func TestLocationAreaInfo(t *testing.T) {
	cases := []struct {
		input    string
		expected LocationAreaInfo
	}{
		{
			input: `{
  "id": 1,
  "name": "canalave-city-area",
  "game_index": 1,
  "encounter_method_rates": [
    {
      "encounter_method": {
        "name": "old-rod",
        "url": "https://pokeapi.co/api/v2/encounter-method/2/"
      },
      "version_details": [
        {
          "rate": 25,
          "version": {
            "name": "platinum",
            "url": "https://pokeapi.co/api/v2/version/14/"
          }
        }
      ]
    }
  ],
  "location": {
    "name": "canalave-city",
    "url": "https://pokeapi.co/api/v2/location/1/"
  },
  "names": [
    {
      "name": "",
      "language": {
        "name": "en",
        "url": "https://pokeapi.co/api/v2/language/9/"
      }
    }
  ],
  "pokemon_encounters": [
    {
      "pokemon": {
        "name": "tentacool",
        "url": "https://pokeapi.co/api/v2/pokemon/72/"
      },
      "version_details": [
        {
          "version": {
            "name": "diamond",
            "url": "https://pokeapi.co/api/v2/version/12/"
          },
          "max_chance": 60,
          "encounter_details": [
            {
              "min_level": 20,
              "max_level": 30,
              "condition_values": [],
              "chance": 60,
              "method": {
                "name": "surf",
                "url": "https://pokeapi.co/api/v2/encounter-method/5/"
              }
            }
          ]
        }
      ]
    }
  ]
}`,
			expected: LocationAreaInfo{
				ID:        1,
				Name:      "canalave-city-area",
				GameIndex: 1,
				Location: Location{
					Name: "canalave-city",
					URL:  "https://pokeapi.co/api/v2/location/1/",
				},
				PokemonEncounters: []PokemonEncounters{
					{
						Pokemon: Pokemon{
							Name: "tentacool",
							URL:  "https://pokeapi.co/api/v2/pokemon/72/",
						},
					},
				},
			},
		},
	}
	for _, c := range cases {
		var locationAreaInfo LocationAreaInfo
		err := json.Unmarshal([]byte(c.input), &locationAreaInfo)
		if err != nil {
			t.Errorf("error trying to unmarshal json: %v", err)
		}
		if c.expected.ID != locationAreaInfo.ID {
			t.Errorf("%v != %v", c.expected.ID, locationAreaInfo.ID)
		}
		if c.expected.Name != locationAreaInfo.Name {
			t.Errorf("%v != %v", c.expected.Name, locationAreaInfo.Name)
		}
		if c.expected.GameIndex != locationAreaInfo.GameIndex {
			t.Errorf("%v != %v", c.expected.GameIndex, locationAreaInfo.GameIndex)
		}
		if c.expected.Location.Name != locationAreaInfo.Location.Name {
			t.Errorf("%v != %v", c.expected.Location.Name, locationAreaInfo.Location.Name)
		}
		if c.expected.Location.URL != locationAreaInfo.Location.URL {
			t.Errorf("%v != %v", c.expected.Location.URL, locationAreaInfo.Location.URL)
		}
		if len(c.expected.PokemonEncounters) != len(locationAreaInfo.PokemonEncounters) {
			t.Errorf("%v != %v", len(c.expected.PokemonEncounters), len(locationAreaInfo.PokemonEncounters))
		}
		for i, encounter := range c.expected.PokemonEncounters {
			if encounter.Pokemon.Name != locationAreaInfo.PokemonEncounters[i].Pokemon.Name {
				t.Errorf("%v != %v", encounter.Pokemon.Name, locationAreaInfo.PokemonEncounters[i].Pokemon.Name)
			}
			if encounter.Pokemon.URL != locationAreaInfo.PokemonEncounters[i].Pokemon.URL {
				t.Errorf("%v != %v", encounter.Pokemon.URL, locationAreaInfo.PokemonEncounters[i].Pokemon.URL)
			}
		}
	}
}

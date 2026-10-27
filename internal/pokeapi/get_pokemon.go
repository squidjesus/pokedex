package pokeapi

import (
	"encoding/json"
	"fmt"
)

func (c Client) PokeReq(name string) (Pokemon, error) {

	// Build the url
	url := baseUrl + "/pokemon/" + name

	// Return Cached data if target URL matches an entry in the cache
	if cached_data, ok := c.cache.Get(url); ok {
		var output Pokemon
		if err := json.Unmarshal(cached_data, &output); err != nil {
			return Pokemon{}, err
		}
		return output, nil
	}

	// Call the ApiRequest for a Pokemon type
	pokemon, err := c.ApiRequest[Pokemon](url)
	if err != nil {
		fmt.Println(err)
		return Pokemon{}, err
	}

	// Return results
	return pokemon, nil
}

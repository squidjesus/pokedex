package pokeapi

import (
	"encoding/json"
)

func (c Client) LocationReqest(pageUrl *string) (LocationResponse, error) {

	// Set url
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	// Return Cached data if target URL matches an entry in the cache
	if cached_data, ok := c.cache.Get(url); ok {
		var output LocationResponse
		if err := json.Unmarshal(cached_data, &output); err != nil {
			return LocationResponse{}, err
		}
		return output, nil
	}

	// Call the ApiRequest for a LocationResponse type
	locations, err := c.ApiRequest[LocationResponse](url)
	if err != nil {
		return LocationResponse{}, err
	}

	// Return results
	return locations, nil
}

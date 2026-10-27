package pokeapi

import (
	"encoding/json"
)

func (c Client) LocDetailReq(areaName string) (LocationDetails, error) {

	url := baseUrl + "/location-area/" + areaName

	// Return Cached data if target URL matches an entry in the cache
	if cached_data, ok := c.cache.Get(url); ok {
		var output LocationDetails
		if err := json.Unmarshal(cached_data, &output); err != nil {
			return LocationDetails{}, err
		}
		return output, nil
	}

	// Call the ApiRequest for a LocationDetails type
	details, err := c.ApiRequest[LocationDetails](url)
	if err != nil {
		return LocationDetails{}, err
	}

	// Return results
	return details, nil
}

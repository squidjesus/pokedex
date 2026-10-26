package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
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

	// Build the http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetails{}, err
	}

	// Make the http request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetails{}, err
	}
	defer res.Body.Close()

	// Read bytes from the http response body
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationDetails{}, err
	}

	// Cache the response as a key (url) value (bytes)
	c.cache.Add(url, bytes)

	// Unmarshal the bytes into a LocationDetails struct
	var output LocationDetails
	if err := json.Unmarshal(bytes, &output); err != nil {
		return LocationDetails{}, err
	}

	return output, nil
}

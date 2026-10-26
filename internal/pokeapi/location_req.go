package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
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

	// Build the http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	// Make the http request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer res.Body.Close()

	// Read bytes from the http response body
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResponse{}, err
	}

	// Cache the response as a key (url) value (bytes)
	c.cache.Add(url, bytes)

	// Unmarshal the bytes into a LocationResponse struct
	var output LocationResponse
	if err := json.Unmarshal(bytes, &output); err != nil {
		return LocationResponse{}, err
	}

	return output, nil
}

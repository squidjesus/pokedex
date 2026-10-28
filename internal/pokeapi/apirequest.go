package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c Client) ApiRequest[T any](url string) (T, error) {

	// Initialize the struct of the given type
	var output T

	// Build the http request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return output, err
	}

	// Make the http request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return output, err
	}
	if res.StatusCode < 200 || res.StatusCode > 399 {
		return output, errors.New("Status code not 200")
	}

	// Read bytes from the http response body
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return output, err
	}
	defer res.Body.Close()

	// Cache the response as a key (url) value (bytes)
	c.cache.Add(url, bytes)

	// Unmarshal the bytes into the struct of our given type
	if err := json.Unmarshal(bytes, &output); err != nil {
		return output, err
	}

	return output, nil
}

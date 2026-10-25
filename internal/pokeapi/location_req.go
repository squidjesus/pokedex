package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c Client) LocationReqest(pageUrl *string) (LocationResponse, error) {

	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer res.Body.Close()

	var output LocationResponse
	if err := json.NewDecoder(res.Body).Decode(&output); err != nil {
		return LocationResponse{}, err
	}

	return output, nil
}

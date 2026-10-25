package pokeApi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type MapResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func MapReq(url string) (MapResponse, error) {

	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return MapResponse{}, fmt.Errorf("Error creating request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return MapResponse{}, fmt.Errorf("Error getting api data: %w", err)
	}
	defer res.Body.Close()

	var output MapResponse
	if err := json.NewDecoder(res.Body).Decode(&output); err != nil {
		return MapResponse{}, fmt.Errorf("Error decoding json: %w", err)
	}

	return output, nil
}

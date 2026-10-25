package main

import (
	"fmt"

	"github.com/squidjesus/pokedex/internal/pokeApi"
)

func commandMap(config *Config) error {

	url := config.mapNextUrl

	response, err := pokeApi.MapReq(url)
	if err != nil {
		return fmt.Errorf("Error calling pokeApi: %w", err)
	}

	//update the 'next' and 'previous' urls
	config.mapNextUrl = response.Next
	if response.Previous != nil {
		config.mapPrevUrl = *response.Previous
	}

	locationAreas := response.Results

	for _, area := range locationAreas {
		fmt.Println(area.Name)
	}

	return nil

}

package main

import (
	"errors"
	"fmt"
)

func commandMapNext(cfg *Config, param *string) error {

	if err := showMap(cfg, cfg.mapNextUrl); err != nil {
		return err
	}

	return nil
}

func commandMapBack(cfg *Config, param *string) error {

	if cfg.mapPrevUrl == nil {
		return errors.New("you're on the first page\n")
	}

	if err := showMap(cfg, cfg.mapPrevUrl); err != nil {
		return err
	}

	return nil
}

func showMap(cfg *Config, url *string) error {

	// call the API
	response, err := cfg.pokeapiClient.LocationReqest(url)
	if err != nil {
		return err
	}

	// update the 'next' and 'previous' urls
	cfg.mapNextUrl = response.Next
	cfg.mapPrevUrl = response.Previous

	// print results
	for _, area := range response.Results {
		fmt.Println(area.Name)
	}

	return nil
}

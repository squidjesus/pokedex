package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *Config, args ...string) error {

	if len(args) < 1 {
		return errors.New("must include a location name")
	}

	areaName := args[0]

	fmt.Printf("Exploring %s...\n", areaName)
	response, err := cfg.pokeapiClient.LocDetailReq(areaName)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range response.PokemonEncounters {

		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}

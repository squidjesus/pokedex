package main

import (
	"fmt"
)

func commandExplore(cfg *Config, areaName *string) error {

	fmt.Printf("Exploring %s...\n", *areaName)
	response, err := cfg.pokeapiClient.LocDetailReq(*areaName)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range response.PokemonEncounters {

		fmt.Println(" - " + encounter.Pokemon.Name)
	}
	return nil
}

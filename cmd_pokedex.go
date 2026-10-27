package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *Config, args ...string) error {
	if len(cfg.caughtPokemon) <= 0 {
		return errors.New("No caught pokemon! Collect some with the catch command.")
	}
	fmt.Println("Your pokedex:")
	for _, caught := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", caught.Name)
	}
	return nil
}

package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return errors.New("requires the name of a (caught) pokemon")
	}
	if pokemon, ok := cfg.caughtPokemon[args[0]]; !ok {
		return errors.New("you have not caught that pokemon")
	} else {
		fmt.Println("Name: " + pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf(" - %s\n", t.Type.Name)
		}
	}
	return nil
}

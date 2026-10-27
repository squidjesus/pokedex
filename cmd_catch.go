package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *Config, args ...string) error {

	if len(args) < 1 {
		return errors.New("must provide a pokemon name")
	}

	pokemon, err := cfg.pokeapiClient.PokeReq(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokemon.Name + "...")

	if catchAttempt(pokemon.BaseExperience) {
		fmt.Println(pokemon.Name + " was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Println(pokemon.Name + " escaped!")
	}
	return err
}

func catchAttempt(baseXp int) bool {
	diffMod := 65
	min := 1
	target := baseXp / diffMod
	if target <= 0 {
		return true
	}
	if baseXp > diffMod {
		target += 1
	}
	roll := rand.IntN(target) + min
	if roll == target {
		return true
	}
	//fmt.Printf("Target is %d, rolled a %d\n", target, roll)
	return false
}

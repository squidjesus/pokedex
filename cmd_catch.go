package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

func commandCatch(cfg *Config, args ...string) error {

	if len(args) < 1 {
		return errors.New("must provide a pokemon name")
	}

	name := args[0]

	pokemon, err := cfg.pokeapiClient.PokeReq(name)
	if err != nil {
		return errors.New(name + " is not a valid pokemon")
	}

	fmt.Printf("Throwing a Pokeball at %s", pokemon.Name)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Printf(".")
	time.Sleep(500 * time.Millisecond)
	fmt.Printf(".")
	time.Sleep(500 * time.Millisecond)

	if catchAttempt(pokemon.BaseExperience) {
		fmt.Println(" " + pokemon.Name + " was caught!")
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Println(" " + pokemon.Name + " escaped!")
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

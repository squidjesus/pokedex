package main

import (
	"time"

	"github.com/squidjesus/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)

	config := &Config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
	}

	startRepl(config)
}

package main

import (
	"time"

	"github.com/squidjesus/pokedex/internal/pokeapi"
	"github.com/squidjesus/pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)

	config := &Config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
		pokeCache:     pokeCache,
	}

	startRepl(config)
}

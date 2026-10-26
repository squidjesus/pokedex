package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/squidjesus/pokedex/internal/pokeapi"
	"github.com/squidjesus/pokedex/internal/pokecache"
)

type Config struct {
	commands      map[string]CliCommand
	pokeapiClient pokeapi.Client
	pokeCache     pokecache.Cache
	mapNextUrl    *string
	mapPrevUrl    *string
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)
	return words
}

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		text := cleanInput(scanner.Text())
		commandName := text[0]
		param := ""
		if len(text) > 1 {
			param = text[1]
		}
		if command, ok := cfg.commands[commandName]; ok {
			err := command.callback(cfg, &param)
			if err != nil {
				fmt.Printf("Error running command %s: %v", command.name, err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

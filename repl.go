package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/squidjesus/pokedex/internal/pokeapi"
)

type Config struct {
	commands      map[string]CliCommand
	pokeapiClient pokeapi.Client
	mapNextUrl    *string
	mapPrevUrl    *string
	caughtPokemon map[string]pokeapi.Pokemon
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
		args := []string{}
		if len(text) > 1 {
			args = text[1:]
		}
		if command, ok := cfg.commands[commandName]; ok {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

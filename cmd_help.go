package main

import (
	"fmt"
)

func commandHelp() error {
	commands := getCommands()
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

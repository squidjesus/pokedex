package main

import (
	"fmt"
)

func commandHelp(config *Config, args ...string) error {
	commands := config.commands
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println()
	for i := range commands {
		fmt.Printf("%s: %s\n", commands[i].name, commands[i].description)
	}
	fmt.Println()
	return nil
}

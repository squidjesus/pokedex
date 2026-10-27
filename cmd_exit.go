package main

import (
	"fmt"
	"os"
)

func commandExit(config *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Closing the Pokedex... Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}

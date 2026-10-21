package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		text := cleanInput(scanner.Text())
		commandName := text[0]

		if command, ok := commands[commandName]; ok {
			err := command.callback()
			if err != nil {
				fmt.Printf("Error running command %s: %v", command.name, err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

}

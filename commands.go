package main

type CliCommand struct {
	name        string
	description string
	callback    func(*Config, *string) error
}

func getCommands() map[string]CliCommand {

	var commands = map[string]CliCommand{

		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},

		"map": {
			name:        "map",
			description: "Display the next 20 location areas",
			callback:    commandMapNext,
		},

		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas",
			callback:    commandMapBack,
		},

		"explore": {
			name:        "explore",
			description: "Lists the pokemon encountered in area NAME, example: 'explore eterna-city-area'",
			callback:    commandExplore,
		},
	}

	return commands
}

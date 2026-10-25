package main

type CliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]CliCommand {

	return map[string]CliCommand{

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
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
	}
}

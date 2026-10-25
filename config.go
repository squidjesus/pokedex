package main

type Config struct {
	commands   map[string]CliCommand
	baseUrl    string
	mapUrl     string
	mapNextUrl string
	mapPrevUrl string
}

func getConfig() *Config {

	config := &Config{
		commands:   getCommands(),
		baseUrl:    "https://pokeapi.co/api/v2/",
		mapNextUrl: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		mapPrevUrl: "",
	}

	return config
}

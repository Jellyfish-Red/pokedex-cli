package main

import "github.com/jellyfish-red/pokedex-cli/internal/pokeapi"

type Command struct {
	name string
	description string
	callback func(*Config) error
}

type Config struct {
	Client pokeapi.APIClient
	Next *string
	Previous *string
}

func getCommands() map[string]Command {
	return map[string]Command {
		"help": {
			name: "help",
			description: "Provides commands usable in this program.",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Close this CLI program.",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Show the next page of locations.",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Show the previous page of locations.",
			callback: commandMapb,
		},
	}
}
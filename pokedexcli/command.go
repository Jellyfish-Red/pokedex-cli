package pokedexcli

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "Provides commands usable in this CLI program.",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Close this CLI program.",
			callback: commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println("Test!")
	return nil
}

func commandExit() error {
	fmt.Println("Closing CLI.")
	os.Exit(0)
	return nil
}
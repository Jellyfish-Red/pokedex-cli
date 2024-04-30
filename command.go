package main

type Command struct {
	name string
	description string
	callback func() error
	// callback func(*Config) error
}

// type Config struct {
// 	Next string
// 	Previous string
// }

func getCommands() map[string]Command {
	return map[string]Command {
		"help": {
			name: "help",
			description: "Provides commands usable in this CLI program.",
			callback: commandHelp,//(&c),
		},
		"exit": {
			name: "exit",
			description: "Close this CLI program.",
			callback: commandExit,//(&c),
		},
		"map": {
			name: "map",
			description: "",
			callback: commandMap,//(&c),
		},
		"mapb": {
			name: "mapb",
			description: "",
			callback: commandMapb,//(&c),
		},
	}
}
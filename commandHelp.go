package main

import "fmt"

func commandHelp(config *Config) error {
	fmt.Println("")
	fmt.Println("List of Commands:")
	fmt.Println("")

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Println("")
	return nil
}
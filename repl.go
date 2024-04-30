package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jellyfish-red/pokedex-cli/internal/pokeapi"
)

func startRepl() {
	client := pokeapi.APIClient{}
	config := &Config {
		Client: client,
	}

	commands := getCommands()
	input := bufio.NewScanner(os.Stdin)

	for {
		// Prepare initial CLI input vector
		fmt.Print("Pokedex > ")

		// Wait for new command
		input.Scan()

		// Sanitize and standardize input
		words := cleanInput(input.Text())
		if len(words) == 0 {
			continue
		}

		// Obtain input and print input back.
		name := words[0]
		// data := words[1:]
		
		if command, ok := commands[name]; ok {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Invalid command: " + name)
			fmt.Println("")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
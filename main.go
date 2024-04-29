package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Pokedex CLI")
	fmt.Println("===========")

	commands := commands.getCommands()

	for {
		// Prepare initial CLI input vector
		input := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter command: ")

		// Wait for new command
		input.Scan()

		// Obtain input and print input back.
		data := input.Text()
		// fmt.Println("Input: ", data)

		if commands[data] {
			commands[data].callback()
		}
	}
}
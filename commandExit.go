package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing CLI.")
	os.Exit(0)
	return nil
}
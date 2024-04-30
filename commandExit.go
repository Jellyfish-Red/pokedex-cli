package main

import (
	"fmt"
	"os"
)

func commandExit(config *Config) error {
	fmt.Println("Closing CLI.")
	os.Exit(0)
	return nil
}
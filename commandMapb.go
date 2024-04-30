package main

import (
	"errors"
	"fmt"
)

func commandMapb(config *Config) error {
	if config.Previous == nil {
		return errors.New("user is already on the first page")
	}

	locations, err := config.Client.requestArea(config.Previous)
	if err != nil {
		return err
	}

	config.Next = locations.Next
	config.Previous = locations.Previous

	for _, location:= range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
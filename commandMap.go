package main

import (
	"fmt"
)

func commandMap(config *Config) error {
	locations, err := config.Client.RequestArea(config.Next)
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
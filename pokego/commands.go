package main

import (
	"errors"
	"fmt"
	"os"

	"pokego/internal/pokeapi"
)

func commandHelp(config *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(config *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(config *Config) error {
	locations, err := pokeapi.PokeApiRequest(config.nextLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locations.Next
	config.prevLocationsURL = locations.Previous

	printLocations(&locations)

	return nil
}

func commandMapb(config *Config) error {
	if config.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locations, err := pokeapi.PokeApiRequest(config.prevLocationsURL)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locations.Next
	config.prevLocationsURL = locations.Previous

	printLocations(&locations)

	return nil
}

func printLocations(res *pokeapi.Response) {
	for _, location := range res.Locations {
		fmt.Println(location.Name)
	}
}

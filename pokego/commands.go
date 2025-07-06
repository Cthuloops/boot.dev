package main

import (
	"errors"
	"fmt"
	"os"

	"pokego/internal/config"
	"pokego/internal/pokeapi"
)

func commandHelp(config *config.Config) error {
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

func commandExit(config *config.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(config *config.Config) error {
	locations, err := pokeapi.PokeApiRequest(config.NextLocationsURL,
		config.Cache)
	if err != nil {
		return err
	}

	config.NextLocationsURL = locations.Next
	config.PrevLocationsURL = locations.Previous

	printLocations(&locations)

	return nil
}

func commandMapb(config *config.Config) error {
	if config.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locations, err := pokeapi.PokeApiRequest(config.PrevLocationsURL,
		config.Cache)
	if err != nil {
		return err
	}

	config.NextLocationsURL = locations.Next
	config.PrevLocationsURL = locations.Previous

	printLocations(&locations)

	return nil
}

func printLocations(res *pokeapi.Response) {
	for _, location := range res.Locations {
		fmt.Println(location.Name)
	}
}

func commandCkey(config *config.Config) error {
	config.Cache.PrintKeys()
	return nil
}

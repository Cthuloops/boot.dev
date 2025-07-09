package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"pokego/internal/config"
	"pokego/internal/pokeapi"
)

func commandHelp(config *config.Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	maxWidth := getMaxCommandNameLength(getCommands())
	for _, cmd := range getCommands() {
		padding := maxWidth - (len(cmd.name) - 1)
		fmt.Printf("%s:%*s%s\n", cmd.name, padding, " ", cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(config *config.Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(config *config.Config, args ...string) error {
	locations, err := config.PokeApiClient.ListLocations(config.NextLocationsURL)
	if err != nil {
		return err
	}

	config.NextLocationsURL = locations.Next
	config.PrevLocationsURL = locations.Previous

	printLocations(&locations)

	return nil
}

func commandMapb(config *config.Config, args ...string) error {
	if config.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locations, err := config.PokeApiClient.ListLocations(config.PrevLocationsURL)
	if err != nil {
		return err
	}

	config.NextLocationsURL = locations.Next
	config.PrevLocationsURL = locations.Previous

	printLocations(&locations)

	return nil
}

func commandExplore(config *config.Config, args ...string) error {
	if len(args) > 1 {
		return errors.New("too many arguments")
	}
	location := pokeapi.GetURL(strings.ToLower(args[0]))
	locationArea, err := config.PokeApiClient.PokemonAtLocation(&location)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locationArea.Name)
	fmt.Println("Found Pokemon:")
	for _, monster := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", monster.Pokemon.Name)
	}
	return nil
}

func printLocations(res *pokeapi.Locations) {
	for _, location := range res.Locations {
		fmt.Println(location.Name)
	}
}

func getMaxCommandNameLength(commands map[string]cliCommand) int {
	maxWidth := 0
	for _, cmd := range commands {
		if len(cmd.name) > maxWidth {
			maxWidth = len(cmd.name)
		}
	}
	return maxWidth
}

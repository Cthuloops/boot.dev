package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"

	"pokego/internal/config"
	"pokego/internal/pokeapi"
)

var (
	ErrTooManyArgs = errors.New("too many arguments")
	ErrTooFewArgs  = errors.New("too few arguments")
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
	if len(args) == 0 {
		return ErrTooFewArgs
	}
	if len(args) > 1 {
		return ErrTooManyArgs
	}
	locationURL := pokeapi.GetLocationURL(strings.ToLower(args[0]))
	locationArea, err := config.PokeApiClient.PokemonAtLocation(&locationURL)
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

func commandCatch(config *config.Config, args ...string) error {
	if len(args) == 0 {
		return ErrTooFewArgs
	}
	if len(args) > 1 {
		return ErrTooManyArgs
	}
	pokename := strings.ToLower(args[0])
	pokeURL := pokeapi.GetPokemonURL(pokename)
	pokemon, err := config.PokeApiClient.PokemonInfo(&pokeURL)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokename)
	fmt.Printf("%s %s\n", pokename, catchPokemon(config, pokemon))
	return nil
}

func commandInspect(config *config.Config, args ...string) error {
	if len(args) == 0 {
		return ErrTooFewArgs
	}
	if len(args) > 1 {
		return ErrTooManyArgs
	}
	pokename := strings.ToLower(args[0])
	pokemon, ok := config.Pokedex[pokename]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	printPokeInfo(pokemon)
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

func catchPokemon(config *config.Config, poke pokeapi.Pokemon) string {
	modifier := math.Log1p(float64(poke.BaseExperience)) / 10.0
	successRate := 0.85 - min(modifier, 0.35)
	roll := rand.Float64()

	if roll >= successRate {
		if _, ok := config.Pokedex[poke.Name]; !ok {
			config.Pokedex[poke.Name] = poke
		}
		return "was caught!"
	}
	return "escaped!"
}

func printPokeInfo(poke pokeapi.Pokemon) {
	fmt.Println("Name:", poke.Name)
	fmt.Println("Height:", poke.Height)
	fmt.Println("Weight:", poke.Weight)
	fmt.Println("Stats:")
	for _, stat := range poke.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, ptype := range poke.Types {
		fmt.Printf("  - %s\n", ptype.Type.Name)
	}
}

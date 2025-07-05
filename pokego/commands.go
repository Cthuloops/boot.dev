package main

import (
	"errors"
	"fmt"
	"os"

	"pokego/internal/pokeapi"
)

func commandHelp(c *pokeapi.Config) error {
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

func commandExit(c *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(c *pokeapi.Config) error {
	if c.GetMapCall() == 0 {
		c.IncMapCall()
		printLocations(c)
		return nil
	}

	if err := c.GetNextPage(); err != nil {
		return err
	}

	printLocations(c)
	return nil
}

func commandMapB(c *pokeapi.Config) error {
	err := c.GetPreviousPage()

	if errors.Is(err, pokeapi.ErrNoPreviousPage) {
		return fmt.Errorf("you're on the first page")
	} else if err != nil {
		return err
	}
	printLocations(c)
	return nil
}

func printLocations(c *pokeapi.Config) {
	for _, loc := range c.Locations {
		fmt.Println(loc.Name)
	}
}

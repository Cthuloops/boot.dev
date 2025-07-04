package main

import (
	"errors"
	"fmt"
	"os"

	"pokego/services"
)

func commandHelp(c *services.Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands(c) {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit(c *services.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(c *services.Config) error {
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

func commandMapB(c *services.Config) error {
	err := c.GetPreviousPage()

	if errors.Is(err, services.ErrNoPreviousPage) {
		return fmt.Errorf("you're on the first page")
	} else if err != nil {
		return err
	}
	printLocations(c)
	return nil
}

func printLocations(c *services.Config) {
	for _, loc := range c.Locations {
		fmt.Println(loc.Name)
	}
}

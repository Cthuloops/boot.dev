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
	if err := c.GetNextPage(); err != nil {
		return err
	}
	for _, loc := range c.Locations {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(c *services.Config) error {
	err := c.GetPreviousPage()

	if errors.Is(err, services.ErrNoPreviousPage) {
		return fmt.Errorf("you're on the first page")
	} else if err != nil {
		return err
	}
	for _, loc := range c.Locations {
		fmt.Println(loc.Name)
	}
	return nil
}

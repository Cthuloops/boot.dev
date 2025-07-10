package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"pokego/internal/config"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)
	config := config.NewConfig(30*time.Second, 5*time.Minute)

	for {
		// Print prompt
		fmt.Print("Pokedex > ")

		reader.Scan()

		// Parse input
		cleanedInput := cleanInput(reader.Text())
		inputLength := len(cleanedInput)
		if inputLength == 0 {
			continue
		}

		// Run the command
		if command, ok := getCommands()[cleanedInput[0]]; ok {
			// Get the args and expand them into strings. If there are any
			// that is.
			err := command.callback(config, cleanedInput[1:]...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	input := strings.ToLower(text)
	cleanedInput := strings.Fields(input)
	return cleanedInput
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config.Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Prints a list of all pokemon at a given location, takes a location as an argument",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the named pokemon",
			callback:    commandCatch,
		},
	}
}

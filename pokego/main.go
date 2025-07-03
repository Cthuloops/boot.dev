package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"pokego/services"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {

	reader := bufio.NewScanner(os.Stdin)
	firstMapCall := true
	for {
		// Print prompt
		fmt.Print("Pokedex > ")

		reader.Scan()

		// Parse input
		cleanedInput := cleanInput(reader.Text())
		if len(cleanedInput) == 0 {
			continue
		}

		// Run the command
		if command, ok := getCommands()[cleanedInput[0]]; ok {
			if (command.name == "map" || command.name == "mapb") && firstMapCall {
			}
			if err := command.callback(); err != nil {
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
			callback:    commandMapB,
		},
	}
}

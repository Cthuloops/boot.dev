package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pokego/services"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*services.Config) error
}

func main() {

	reader := bufio.NewScanner(os.Stdin)
	firstMapCall := true
	config, err := services.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

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
		if command, ok := getCommands(config)[cleanedInput[0]]; ok {
			// I feel like there's a better way to do this.
			if command.name == "map" && firstMapCall {
				for _, loc := range config.Locations {
					fmt.Println(loc.Name)
				}
				firstMapCall = false
			}
			if err := command.callback(config); err != nil {
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

func getCommands(c *services.Config) map[string]cliCommand {
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

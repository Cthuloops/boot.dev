package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var cliCommands map[string]cliCommand

func main() {
	cliCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Help Message",
			callback:    commandHelp,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		inputScanner := bufio.NewScanner(os.Stdin)
		var input string
		if inputScanner.Scan() {
			input += inputScanner.Text()
		}
		cleanedInput := cleanInput(input)
		fmt.Printf("Your command was: %s\n", cleanedInput[0])
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Usage:\n\n")
	for _, command := range cliCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func cleanInput(text string) []string {
	if text == ""
}
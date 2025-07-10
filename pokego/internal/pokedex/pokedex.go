package pokedex

import (
	"fmt"

	"pokego/internal/pokeapi"
)

type Pokedex map[string]pokeapi.Pokemon

func (p Pokedex) PrintEntryNames() {
	fmt.Println("Your Pokedex:")
	for name := range p {
		fmt.Printf(" - %s\n", name)
	}
}

package config

import (
	"time"

	"pokego/internal/pokeapi"
	"pokego/internal/pokedex"
)

type Config struct {
	PokeApiClient    pokeapi.Client
	Pokedex          pokedex.Pokedex
	NextLocationsURL *string
	PrevLocationsURL *string
}

func NewConfig(timeout, interval time.Duration) *Config {
	config := Config{
		PokeApiClient: pokeapi.NewClient(timeout, interval),
		Pokedex:       pokedex.Pokedex{},
	}
	return &config
}

package config

import (
	"time"

	"pokego/internal/pokeapi"
)

type Config struct {
	PokeApiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
}

func NewConfig(timeout, interval time.Duration) *Config {
	config := Config{
		PokeApiClient: pokeapi.NewClient(timeout, interval),
	}
	return &config
}

package config

import (
	"time"

	"pokego/internal/pokeapi"
)

type Config struct {
	pokeapiClient    pokeapi.Client
	NextLocationsURL *string
	PrevLocationsURL *string
}

func NewConfig(timeout, interval time.Duration) *Config {
	config := Config{
		pokeapiClient: pokeapi.NewClient(timeout, interval),
	}
	return &config
}

package config

import (
	"time"

	"pokego/internal/pokecache"
)

type Config struct {
	NextLocationsURL *string
	PrevLocationsURL *string
	Cache            *pokecache.Cache
}

func NewConfig(interval time.Duration) *Config {
	var config Config
	config.Cache = pokecache.NewCache(interval)
	return &config
}
